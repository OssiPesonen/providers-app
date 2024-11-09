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
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Server.Port))

	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	auth, err := auth.New(config.Auth.JWTSecret)
	if err != nil {
		logger.Fatalf("failed to init auth service: %v", err)
	}

	// Authentication middleware. Skip Login and Register methods
	authInterceptor, err := interceptor.NewAuthInterceptor(auth, logger, []string{"Register", "Login"})
	if err != nil {
		logger.Fatalf("failed to create auth interceptor: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.UnaryAuthMiddleware),
	)

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

	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

	return s
}

type ServerError struct {
	Message string
	Code    codes.Code
}

func (s *Server) FromError(err error) ServerError {
	var serverError ServerError
	var serviceError core.Error

	if errors.As(err, &serviceError) {
		serverError.Message = serviceError.ApplicationError().Error()
		serviceError := serviceError.ServiceError()

		switch serviceError {
		case core.ErrRevokedRefreshToken:
		case core.ErrExpiredRefreshToken:
			serverError.Code = codes.Unauthenticated
		case core.ErrUserNotFound:
			serverError.Code = codes.Unauthenticated
		case core.ErrUserAlreadyExists:
			serverError.Code = codes.AlreadyExists
		case core.ErrInvalidPassword:
			serverError.Code = codes.Unauthenticated

		default:
			// Uncontrolled and unidentified error types are internal errors
			serverError.Code = codes.Internal
		}
	}

	return serverError
}
