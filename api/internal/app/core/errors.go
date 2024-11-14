package core

import "errors"

var (
	// Refresh token was used to gain new access token, but it is already revoked
	ErrRevokedRefreshToken = errors.New("refresh token is revoked")
	// Refresh token was used to gain new access token, but it is expired
	ErrExpiredRefreshToken = errors.New("refresh token has expired")
	// Error cannot be safely interpreted
	ErrInternal = errors.New("internal server error")
	// Entity / Resource is not found
	ErrNotFound = errors.New("resource not found")
	// Someone attempted to register a user account,
	//but the user (with email) already exists in the system
	ErrUserAlreadyExists = errors.New("user already exists")
	// Someone attempted to authenticate with an email that is not found
	ErrUserNotFound = errors.New("user not found")
	// Name + City collision
	ErrProviderAlreadyExists = errors.New("a provider with this name already exists for the city")
	// User entered an invalid password
	ErrInvalidPassword = errors.New("invalid password")
)

type Error struct {
	// Application level error which must not be outputted
	applicationError error
	// Service level error ie. controlled
	serviceError error
}

func NewError(serviceError, applicationError error) error {
	return Error{
		serviceError:     serviceError,
		applicationError: applicationError,
	}
}

func (e Error) Error() string {
	return errors.Join(e.serviceError, e.applicationError).Error()
}
func (e Error) ApplicationError() error {
	return e.applicationError
}

func (e Error) ServiceError() error {
	return e.serviceError
}
