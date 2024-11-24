package providers

import (
	"log"

	"github.com/ossipesonen/providers-app/internal/app/core"
	"github.com/ossipesonen/providers-app/internal/app/core/models"
	"github.com/upper/db/v4"
)

// Define repository interface that this service needs
type IProviderRepository interface {
	List() (*[]models.Provider, error)
	listForUser(userId int) (*[]models.Provider, error)
	Read(id int) (*models.Provider, error)
	Find(name string, city string) (*models.Provider, error)
	Create(*models.Provider) (int, error)
}

type ProviderService struct {
	logger     *log.Logger
	repository IProviderRepository
}

// This is the entrypoint to the module
func NewProviderService(repository IProviderRepository, logger *log.Logger) *ProviderService {
	return &ProviderService{
		logger:     logger,
		repository: repository,
	}
}

func (s *ProviderService) ListProviders() (*[]models.Provider, error) {
	providers, err := s.repository.List()

	if err != nil {
		s.logger.Printf("Fetching providers failed: %v", err)
		return nil, err
	}

	return providers, nil
}

func (s *ProviderService) ListProvidersForUser(userId int) (*[]models.Provider, error) {
	providers, err := s.repository.listForUser(userId)
	if err != nil {
		s.logger.Printf("Fetching providers for user failed: %v", err)
		return nil, err
	}

	return providers, nil
}

func (s *ProviderService) GetProvider(id int) (*models.Provider, error) {
	provider, err := s.repository.Read(id)

	if err != nil {
		if err != db.ErrNoMoreRows {
			s.logger.Printf("Fetching providers from storage failed:: %v", err)
		}

		return nil, err
	}

	return provider, nil
}

func (s *ProviderService) CreateProvider(provider *models.Provider) (int, error) {
	// Prevent a provider with the same name, in the same city, to exist in the system
	existingProvider, err := s.repository.Find(provider.Name, provider.City)
	if existingProvider != nil {
		return 0, core.NewError(core.ErrProviderAlreadyExists, nil)
	}
	if err != nil {
		s.logger.Printf("An error occurred when attempting to find provider: %v", err)
		return 0, core.NewError(core.ErrInternal, err)
	}

	id, err := s.repository.Create(provider)
	if err != nil {
		s.logger.Printf("An error occurred when attempting to create provider: %v", err)
		return 0, core.NewError(core.ErrInternal, err)
	}

	return id, nil
}
