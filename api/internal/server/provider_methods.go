package server

import (
	"context"
	"database/sql"

	pb "github.com/ossipesonen/go-traffic-lights/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// List all commands available for server
func (s *Server) ListProviders(context context.Context, _ *emptypb.Empty) (*pb.ListProviderResponse, error) {
	providers, err := s.App.Services.Provider.ListProviders()

	if err != nil {
		s.Logger.Printf("Fetching providers failed: %v", err)
		return nil, err
	}

	providersSlice := []*pb.ReadProviderResponse{}

	for _, result := range *providers {
		p := pb.ReadProviderResponse{
			Id:   int32(result.Id),
			Name: result.Name,
		}

		providersSlice = append(providersSlice, &p)
	}

	return &pb.ListProviderResponse{
		Providers: providersSlice,
	}, nil
}

func (s *Server) ReadProvider(ctx context.Context, in *pb.ReadProviderRequest) (*pb.ReadProviderResponse, error) {
	id := in.Id
	provider, err := s.App.Services.Provider.GetProvider(int(id))

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "")
		} else {
			// Log server error
			s.Logger.Printf("Fetching providers from storage failed:: %v", err)
		}

		// Server error
		return nil, status.Error(codes.Internal, "Something went wrong")
	}

	return &pb.ReadProviderResponse{
		Id:   int32(provider.Id),
		Name: provider.Name,
	}, nil
}
