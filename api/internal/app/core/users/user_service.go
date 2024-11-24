package users

import (
	"errors"
	"log"
	"time"

	"github.com/ossipesonen/providers-app/internal/app/auth"
	"github.com/ossipesonen/providers-app/internal/app/core"
	"github.com/ossipesonen/providers-app/internal/app/core/models"
	"github.com/upper/db/v4"
)

// Define repository interface that this service needs
type IUserRepository interface {
	Read(id int) (*models.User, error)
	Find(email string) (*models.User, error)
	Add(user *models.User) error
	SaveRefreshToken(refreshTokenEntry *models.RefreshTokenEntry) error
	GetRefreshToken(refreshToken string) (*models.RefreshTokenEntry, error)
	RevokeRefreshToken(refreshToken string) error
	RevokeAllRefreshTokens(userId int) error
}

type UserService struct {
	logger     *log.Logger
	repository IUserRepository
	auth       *auth.Auth
}

func NewUserService(repository IUserRepository, auth *auth.Auth, logger *log.Logger) *UserService {
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

	// User not found
	if err != nil {
		return nil, core.NewError(core.ErrUserNotFound, err)
	}

	err = s.auth.Password.Compare(user.Password, user.Salt, password)

	// Password comparison failed -> unauthenticated
	if err != nil {
		return nil, core.NewError(core.ErrInvalidPassword, err)
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
		return nil, core.NewError(core.ErrUserAlreadyExists, errors.New("user already exists"))
	}

	hashSalt, err := s.auth.Password.GenerateHash([]byte(userInfo.Password), []byte{})

	if err != nil {
		return nil, core.NewError(core.ErrInternal, err)
	}

	user := models.User{
		Username: userInfo.Username,
		Email:    userInfo.Email,
		Password: string(hashSalt.Hash),
		Salt:     string(hashSalt.Salt),
	}

	err = s.repository.Add(&user)

	if err != nil {
		return nil, core.NewError(core.ErrInternal, err)
	}

	return &user, nil
}

func (s *UserService) Find(userId int) (*models.User, error) {
	user, err := s.repository.Read(userId)

	if err != nil {
		return nil, core.NewError(core.ErrInternal, err)
	}

	return user, nil
}

func (s *UserService) GenerateTokens(user *models.User) (*auth.IssuedTokens, error) {
	tokens, err := s.auth.IssueToken(user.Id)
	if err != nil {
		return nil, core.NewError(core.ErrInternal, err)
	}

	// Refresh token is persisted in storage so we can revoke it as it's used to refresh
	err = s.repository.SaveRefreshToken(&models.RefreshTokenEntry{
		RefreshToken: tokens.RefreshToken,
		Expires:      time.Now().Add(time.Hour * 24).UTC(),
		UserId:       user.Id,
	})

	if err != nil {
		return nil, core.NewError(core.ErrInternal, err)
	}

	return &auth.IssuedTokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (s *UserService) RefreshTokens(refreshToken string) (*auth.IssuedTokens, error) {
	// Ensure token is still valid
	token, err := s.repository.GetRefreshToken(refreshToken)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return nil, core.NewError(core.ErrRevokedRefreshToken, err)
		}

		return nil, core.NewError(core.ErrInternal, err)
	}

	if token.Expires.Before(time.Now()) {
		// Delete the refresh token as it is already expired
		go s.RevokeRefreshToken(refreshToken)
		return nil, core.NewError(core.ErrExpiredRefreshToken, err)
	}

	user, err := s.repository.Read(token.UserId)
	if err != nil {
		if err == db.ErrNoMoreRows {
			// This means user is no longer in the system so it's am edge case
			return nil, core.NewError(core.ErrRevokedRefreshToken, err)
		}

		return nil, core.NewError(core.ErrUserNotFound, err)
	}

	// Generate new pair
	tokens, err := s.GenerateTokens(user)
	if err != nil {
		return nil, err
	}

	// Finally revoke used refresh token
	go s.RevokeRefreshToken(refreshToken)

	return &auth.IssuedTokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (s *UserService) RevokeRefreshToken(refreshToken string) error {
	// Check that the refresh token exists
	_, err := s.repository.GetRefreshToken(refreshToken)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return core.NewError(core.ErrNotFound, err)
		}

		return core.NewError(core.ErrInternal, err)
	}

	err = s.repository.RevokeRefreshToken(refreshToken)

	if err != nil {
		// log here as we also run this call on a goroutine to revoke refresh tokens from storage
		s.logger.Printf("something went wrong when attempting to revoke refresh token: %v", err)
		return core.NewError(core.ErrInternal, err)
	}

	return nil
}

func (s *UserService) RevokeAllRefreshTokens(userId int) error {
	err := s.repository.RevokeAllRefreshTokens(userId)

	if err != nil {
		return core.NewError(core.ErrInternal, err)
	}

	return nil
}
