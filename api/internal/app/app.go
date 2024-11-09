package app

import (
	"log"

	"github.com/ossipesonen/go-traffic-lights/internal/app/auth"
	"github.com/ossipesonen/go-traffic-lights/internal/app/core/models"
	"github.com/ossipesonen/go-traffic-lights/internal/app/core/repositories"
	"github.com/ossipesonen/go-traffic-lights/internal/app/core/services"
	"github.com/ossipesonen/go-traffic-lights/internal/config"
	"github.com/ossipesonen/go-traffic-lights/pkg/database"
)

// Define the API to components
type UserService interface {
	// Authenticates a user using provided email and password.
	// Returns the user resource if successful
	Authenticate(email string, password string) (*models.User, error)
	// Creates a new user resource
	// Returns the resource
	CreateUser(userInfo *models.UserInfo) (*models.User, error)
	// Finds a user resource using only the ID
	Find(userId int) (*models.User, error)
	// Generate access and refresh token for user
	GenerateTokens(user *models.User) (*auth.IssuedTokens, error)
	// Generate new pair of access and refresh token for user
	RefreshTokens(refreshToken string, userId int) (*auth.IssuedTokens, error)
	// Revoke a specific refresh token for given user
	RevokeRefreshToken(refreshToken string, userId int) error
	// Revoke all active refresh tokens for user
	RevokeAllRefreshTokens(userId int) error
}

type ProviderService interface {
	// Lists all providers
	ListProviders() (*[]models.Provider, error)
	// Fetch a single provider using an identifier
	// Returns the provider or error. Error can be sql.ErrnoRows,
	// which indicates no resource found
	GetProvider(id int) (*models.Provider, error)
}

type Services struct {
	User     UserService
	Provider ProviderService
}

type App struct {
	Services *Services
}

// Bootstrap application
// Setup repositories and services to be used by Server methods
func New(config *config.Config, db database.Database, logger *log.Logger) *App {
	srvc := &Services{}

	auth, err := auth.New(config.Auth.JWTSecret)

	if err != nil {
		logger.Fatalf("unable to init auth service :%v", err)
	}

	// Repositores should have db as a dependency.
	// Repositories should not be directly accessible.
	userRepository := repositories.NewUserRepository(db, logger)
	providerRepository := repositories.NewProviderRepository(db, logger)

	// Services should have repository as a depency if data access is required
	// Service is the entry point to a component.
	srvc.User = services.NewUserService(userRepository, auth, logger)
	srvc.Provider = services.NewProviderService(providerRepository, logger)

	return &App{
		Services: srvc,
	}
}
