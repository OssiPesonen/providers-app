package server

import (
	"context"
	"strconv"

	"github.com/ossipesonen/go-traffic-lights/internal/app/core/models"
	"github.com/ossipesonen/go-traffic-lights/internal/server/interceptor"
	pb "github.com/ossipesonen/go-traffic-lights/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// List all commands available for server
func (s *Server) GetToken(ctx context.Context, in *pb.LoginRequest) (*pb.TokenResponse, error) {
	user, err := s.App.Services.User.Authenticate(in.Email, in.Password)

	if err != nil {
		e := s.FromError(err)
		// Determine the error code, but do not output a message here
		return nil, status.Error(e.Code, "")
	}

	tokens, err := s.App.Services.User.GenerateTokens(user)

	if err != nil {
		e := s.FromError(err)
		return nil, status.Error(e.Code, "")
	}

	return &pb.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: &tokens.RefreshToken,
		Exp:          3600,
		TokenType:    "Bearer",
	}, nil
}

func (s *Server) RegisterUser(ctx context.Context, in *pb.RegistrationRequest) (*emptypb.Empty, error) {
	_, err := s.App.Services.User.CreateUser(&models.UserInfo{
		Email:    in.Email,
		Password: in.Password,
		Username: in.Username,
	})

	if err != nil {
		e := s.FromError(err)

		// User already exists. Use code 200 to indicate that
		// request was successful but no resource was created.
		// Do not tell caller that this resource does not exist!
		if e.Code == codes.AlreadyExists {
			return &emptypb.Empty{}, nil
		}

		s.Logger.Printf("An internal server error occurred: %v", err)
		return nil, status.Error(codes.Internal, "")
	}

	// This should be replaced with a 200 OK response and something like
	// email verification being sent, but we shall not do that now
	return &emptypb.Empty{}, nil
}

func (s *Server) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.TokenResponse, error) {
	userId := ctx.Value(interceptor.UserIdKey).(string)

	if userId == "" {
		return nil, status.Error(codes.InvalidArgument, "user id is missing")
	}

	userIdInt, _ := strconv.Atoi(userId)
	tokens, err := s.App.Services.User.RefreshTokens(in.RefreshToken, userIdInt)

	if err != nil {
		s.Logger.Printf("something went wrong when attempting to revoke refresh token: %v", err)
		e := s.FromError(err)
		return nil, status.Error(e.Code, e.Message)
	}

	return &pb.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: &tokens.RefreshToken,
		Exp:          3600,
		TokenType:    "Bearer",
	}, nil
}

func (s *Server) RevokeRefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*emptypb.Empty, error) {
	userId := ctx.Value(interceptor.UserIdKey).(string)
	if userId == "" {
		return nil, status.Error(codes.FailedPrecondition, "user ID was not determined")
	}

	userIdInt, _ := strconv.Atoi(userId)
	err := s.App.Services.User.RevokeRefreshToken(in.RefreshToken, userIdInt)
	if err != nil {
		e := s.FromError(err)
		return nil, status.Error(e.Code, e.Message)
	}

	return &emptypb.Empty{}, nil
}

// Method to allow users to revoke all refresh tokens that are active
func (s *Server) RevokeAllRefreshTokens(ctx context.Context) (*emptypb.Empty, error) {
	userId := ctx.Value(interceptor.UserIdKey).(string)
	if userId == "" {
		return nil, status.Error(codes.FailedPrecondition, "user ID was not determined")
	}

	userIdInt, _ := strconv.Atoi(userId)
	err := s.App.Services.User.RevokeAllRefreshTokens(userIdInt)

	if err != nil {
		e := s.FromError(err)
		return nil, status.Error(e.Code, e.Message)
	}

	return &emptypb.Empty{}, nil
}
