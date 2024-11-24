package server

import (
	"context"
	"strconv"

	"github.com/ossipesonen/providers-app/internal/app/core/models"
	"github.com/ossipesonen/providers-app/internal/server/interceptor"
	pb "github.com/ossipesonen/providers-app/proto"
	"github.com/upper/db/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// List all commands available for server
func (s *Server) ListProviders(context context.Context, _ *emptypb.Empty) (*pb.ListOfProviders, error) {
	providers, err := s.App.Services.Provider.ListProviders()

	if err != nil {
		s.Logger.Printf("Fetching providers failed: %v", err)
		return nil, err
	}

	providersSlice := []*pb.Provider{}
	for _, result := range *providers {
		p := pb.Provider{
			Id:             int32(result.Id),
			Name:           result.Name,
			Region:         result.Region,
			City:           result.City,
			LineOfBusiness: result.LineOfBusiness,
		}

		providersSlice = append(providersSlice, &p)
	}

	return &pb.ListOfProviders{
		Providers: providersSlice,
	}, nil
}

func (s *Server) ReadProvider(ctx context.Context, in *pb.ReadProviderRequest) (*pb.Provider, error) {
	id := in.Id
	provider, err := s.App.Services.Provider.GetProvider(int(id))

	if err != nil {
		if err == db.ErrNoMoreRows {
			return nil, status.Error(codes.NotFound, "")
		} else {
			// Log server error
			s.Logger.Printf("Fetching providers from storage failed:: %v", err)
		}

		// Server error
		return nil, status.Error(codes.Internal, "Something went wrong")
	}

	return &pb.Provider{
		Id:   int32(provider.Id),
		Name: provider.Name,
	}, nil
}

func (s *Server) CreateProvider(ctx context.Context, in *pb.CreateProviderRequest) (*pb.ProviderId, error) {
	userId := ctx.Value(interceptor.UserIdKey).(string)
	if userId == "" {
		return nil, status.Error(codes.FailedPrecondition, "user ID was not determined")
	}

	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		return nil, status.Error(codes.Internal, "Something went wrong. Please try again later.")
	}

	providerId, err := s.App.Services.Provider.CreateProvider(&models.Provider{
		Name:   in.Name,
		City:   in.City,
		Region: in.Region,
		UserId: intUserId,
	})

	if err != nil {
		e := s.FromError(err)
		return nil, status.Error(e.Code, e.Message)
	}

	return &pb.ProviderId{Id: int32(providerId)}, nil
}
