package server

import (
	"context"
	"database/sql"

	"github.com/ossipesonen/go-traffic-lights/internal/app/core/models"
	pb "github.com/ossipesonen/go-traffic-lights/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// List all commands available for server
func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.App.Services.User.Authenticate(in.Email, in.Password)

	if err != nil {
		if err != sql.ErrNoRows {
			s.Logger.Printf("Something went wrong when fetching user: %v", err)
		}

		errUnauthenticated := status.Error(codes.Unauthenticated, "")
		return nil, errUnauthenticated
	}

	tokens, err := s.App.Services.User.GenerateTokens(user)

	if err != nil {
		s.Logger.Printf("Something went wrong when generating tokens: %v", err)
		return nil, status.Error(codes.Internal, "Something went wrong. Please try again later.")
	}

	return &pb.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: &tokens.RefreshToken,
		Exp:          3600,
		TokenType:    "Bearer",
	}, nil
}

func (s *Server) Register(ctx context.Context, in *pb.RegistrationRequest) (*emptypb.Empty, error) {
	_, err := s.App.Services.User.CreateUser(&models.UserInfo{
		Email:    in.Email,
		Password: in.Password,
		Username: in.Username,
	})

	if err != nil {
		// User already exists. Use code 200 to indicate that
		// request was successful but no resource was created
		if err.Error() == "user-already-exists" {
			return &emptypb.Empty{}, nil
		}

		s.Logger.Printf("An internal server error occurred: %v", err)
		return nil, status.Error(codes.Internal, "")
	}

	// This should be replaced with a 200 OK response and
	// email verification being sent, but we shall not do that now
	return &emptypb.Empty{}, nil
}
