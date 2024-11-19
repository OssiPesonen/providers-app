package server

import (
	"errors"
	"fmt"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ossipesonen/go-traffic-lights/internal/app"
	"github.com/ossipesonen/go-traffic-lights/internal/app/auth"
	"github.com/ossipesonen/go-traffic-lights/internal/app/core"
	"github.com/ossipesonen/go-traffic-lights/internal/config"
	"github.com/ossipesonen/go-traffic-lights/internal/server/interceptor"
	"github.com/ossipesonen/go-traffic-lights/pkg/database"
	pb "github.com/ossipesonen/go-traffic-lights/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// Server must implement TrafficLightsServiceClient interface
type Server struct {
	pb.UnimplementedTrafficLightsServiceServer
	Logger *log.Logger
	App    *app.App
	Auth   *auth.Auth
}

// Set up gRPC server to listen for method calls
func New(config *config.Config, logger *log.Logger) *grpc.Server {
	logger.Printf("server starting on port: %d", config.Server.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Server.Port))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	auth, err := auth.New(config.Auth.JWTSecret)
	if err != nil {
		logger.Fatalf("failed to init auth service: %v", err)
	}

	// Authentication middleware. Skip Login and Register methods
	authInterceptor, err := interceptor.NewAuthInterceptor(auth, logger, []string{"GetToken", "RefreshToken", "ListProviders", "RegisterUser"})
	if err != nil {
		logger.Fatalf("failed to create auth interceptor: %v", err)
	}

	s := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(authInterceptor.UnaryAuthMiddleware),
	)

	reflection.Register(s)

	// Setup App to call from methods
	db := database.New(&database.DBConfig{
		Username: config.DB.Username,
		Password: config.DB.Password,
		Host:     config.DB.Host,
		Port:     config.DB.Port,
		Database: config.DB.Database,
		Schema:   config.DB.Schema,
	}, logger)

	app := app.New(config, db, logger)

	pb.RegisterTrafficLightsServiceServer(s, &Server{
		Logger: logger,
		App:    app,
		Auth:   auth,
	})

	// start gRPC server
	log.Println("starting gRPC server...")

	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

	return s
}

type ServerError struct {
	Message string
	Code    codes.Code
}

// Map errors from core services to server that gets outputted to client
// we want to hide certain errors to not expose certain details
func (s *Server) FromError(err error) ServerError {
	var serverError ServerError
	var serviceError core.Error

	if errors.As(err, &serviceError) {
		serviceError := serviceError.ServiceError()
		serverError.Message = serviceError.Error()

		switch serviceError {
		case core.ErrRevokedRefreshToken:
			serverError.Code = codes.Unauthenticated
		case core.ErrExpiredRefreshToken:
			serverError.Code = codes.Unauthenticated
		case core.ErrUserNotFound:
			serverError.Code = codes.Unauthenticated
			// Do not expose message to callers
			serverError.Message = ""
		case core.ErrUserAlreadyExists:
			serverError.Code = codes.AlreadyExists
			// Do not expose message to callers
			serverError.Message = ""
		case core.ErrProviderAlreadyExists:
			serverError.Code = codes.AlreadyExists
		case core.ErrNotFound:
			serverError.Code = codes.NotFound
		case core.ErrInvalidPassword:
			serverError.Code = codes.Unauthenticated
			// Do not expose message to callers
			serverError.Message = ""

		default:
			// Uncontrolled and unidentified error types are internal errors
			serverError.Code = codes.Internal
		}
	}

	return serverError
}
