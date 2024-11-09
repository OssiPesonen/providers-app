package core

import "errors"

var (
	ErrRevokedRefreshToken = errors.New("refresh token is revoked")
	ErrExpiredRefreshToken = errors.New("refresh token has expired")
	ErrInternal            = errors.New("internal server error")
	ErrNotFound            = errors.New("resource not found")
	ErrUserAlreadyExists   = errors.New("")
	// User not found. Do not expose that to caller.
	ErrUserNotFound = errors.New("")
	// Invalid credentials. Do not expose that to caller.
	ErrInvalidPassword = errors.New("")
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
