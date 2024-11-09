package services

import (
	"database/sql"
	"log"

	"github.com/ossipesonen/go-traffic-lights/internal/app/core/models"
)

// Define repository interface that this service needs
type ProviderRepository interface {
	List() (*[]models.Provider, error)
	Read(id int) (*models.Provider, error)
}

type ProviderService struct {
	Logger     *log.Logger
	Repository ProviderRepository
}

// This is the entrypoint to the module
func NewProviderService(repository ProviderRepository, logger *log.Logger) *ProviderService {
	return &ProviderService{
		Logger:     logger,
		Repository: repository,
	}
}

func (s *ProviderService) ListProviders() (*[]models.Provider, error) {
	providers, err := s.Repository.List()

	if err != nil {
		s.Logger.Printf("Fetching providers failed: %v", err)
		return nil, err
	}

	return providers, nil
}

func (s *ProviderService) GetProvider(id int) (*models.Provider, error) {
	provider, err := s.Repository.Read(id)

	if err != nil {
		if err != sql.ErrNoRows {
			s.Logger.Printf("Fetching providers from storage failed:: %v", err)
		}

		return nil, err
	}

	return provider, nil
}
