package services

import (
	"errors"
	"log"
	"time"

	"github.com/ossipesonen/go-traffic-lights/internal/app/auth"
	"github.com/ossipesonen/go-traffic-lights/internal/app/core/models"
)

// Define repository interface that this service needs
type UserRepository interface {
	Read(id int) (*models.User, error)
	Find(email string) (*models.User, error)
	Add(user *models.User) error
	SaveRefreshToken(refreshTokenEntry *models.RefreshTokenEntry) error
	RevokeRefreshToken(userId int) error
}

type UserService struct {
	logger     *log.Logger
	repository UserRepository
	auth       *auth.Auth
}

func NewUserService(repository UserRepository, auth *auth.Auth, logger *log.Logger) *UserService {
	return &UserService{
		logger:     logger,
		repository: repository,
		auth:       auth,
	}
}

// Authenticate user by first looking them up, then comparing
// the provided password with stored hash.
// Returns User if successfully authenticated, otherwise error
func (s *UserService) Authenticate(email string, password string) (*models.User, error) {
	user, err := s.repository.Find(email)

	// Not found
	if err != nil {
		return nil, err
	}

	err = s.auth.Password.Compare(user.Password, user.Salt, password)

	// Unauthenticated
	if err != nil {
		return nil, err
	}

	// User authenticated
	return user, nil
}

// Creates a new user resource
// Returns a pointer to the User if successful, error otherwise
func (s *UserService) CreateUser(userInfo *models.UserInfo) (*models.User, error) {
	// Attempt to find an existing user
	existingUser, _ := s.repository.Find(userInfo.Email)

	if existingUser != nil {
		return nil, errors.New("user-already-exists")
	}

	hashSalt, err := s.auth.Password.GenerateHash([]byte(userInfo.Password), []byte{})

	if err != nil {
		s.logger.Printf("hashing password failed: %v", err)
		return nil, err
	}

	user := models.User{
		Username: userInfo.Username,
		Email:    userInfo.Email,
		Password: string(hashSalt.Hash),
		Salt:     string(hashSalt.Salt),
	}

	err = s.repository.Add(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) Find(userId int) (*models.User, error) {
	user, err := s.repository.Read(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GenerateTokens(user *models.User) (*auth.IssuedTokens, error) {
	tokens, err := s.auth.IssueToken(user.Id)
	if err != nil {
		s.logger.Printf("Unable to generate tokens for user: %v", err)
		return nil, errors.New("unable to issue tokens")
	}

	// Refresh token is persisted in storage so we can revoke it as it's used to refresh
	err = s.repository.SaveRefreshToken(&models.RefreshTokenEntry{
		RefreshToken: tokens.RefreshToken,
		Expires:      time.Now().Add(time.Hour * 24),
		UserId:       user.Id,
	})

	if err != nil {
		s.logger.Printf("Unable to persist refresh token: %v", err)
		return nil, errors.New("unable to save refresh token")
	}

	return &auth.IssuedTokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
