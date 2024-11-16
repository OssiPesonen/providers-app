package interceptor

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Empty struct{}
type MockValidator struct{}

func NewMockValidator() *MockValidator {
	return &MockValidator{}
}

func (m *MockValidator) ValidateToken(token string) (string, error) {
	if token != "valid-token" {
		return "", errors.New("invalid token")
	}
	return "1234", nil
}

var logger = log.New(os.Stdout, "http: ", log.LstdFlags)
var req = &Empty{}
var validator = NewMockValidator()

func TestNewAuthInterceptor(t *testing.T) {
	_, err := NewAuthInterceptor(nil, logger, []string{})
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "validator cannot be nil")
}

func TestUnarityAuthMiddlewareWithoutMetadata(t *testing.T) {
	var ctx = context.Background()
	info := &grpc.UnaryServerInfo{
		FullMethod: "hello",
	}

	mockHandlerCallCounter := 0
	mockHandler := func(ctx context.Context, req any) (any, error) {
		mockHandlerCallCounter++
		return nil, nil
	}

	interceptor, err := NewAuthInterceptor(validator, logger, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, interceptor)

	_, err = interceptor.UnaryAuthMiddleware(ctx, req, info, mockHandler)

	assert.NotNil(t, err)
	assert.Equal(t, "rpc error: code = Unauthenticated desc = Metadata has not been provided", err.Error())
	assert.Equal(t, 0, mockHandlerCallCounter)
}

func TestUnarityAuthMiddlewareWithWhitelistedRoute(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.Background(), map[string][]string{})

	info := &grpc.UnaryServerInfo{
		FullMethod: "hello-world",
	}

	mockHandlerCallCounter := 0
	mockHandler := func(ctx context.Context, req any) (any, error) {
		mockHandlerCallCounter++
		return "foo-bar", nil
	}

	interceptor, err := NewAuthInterceptor(validator, logger, []string{"hello-world"})
	assert.Nil(t, err)
	assert.NotNil(t, interceptor)

	// Should call handler without token
	response, err := interceptor.UnaryAuthMiddleware(ctx, req, info, mockHandler)

	assert.Nil(t, err)
	assert.Equal(t, "foo-bar", response)
	assert.Equal(t, 1, mockHandlerCallCounter)
}

func TestUnarityAuthMiddlewareNoAccessToken(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.Background(), map[string][]string{})
	info := &grpc.UnaryServerInfo{
		FullMethod: "foo-bar",
	}

	mockHandlerCallCounter := 0
	mockHandler := func(ctx context.Context, req any) (any, error) {
		mockHandlerCallCounter++
		return "foo-bar", nil
	}

	interceptor, err := NewAuthInterceptor(validator, logger, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, interceptor)

	// Should call handler
	response, err := interceptor.UnaryAuthMiddleware(ctx, req, info, mockHandler)

	assert.NotNil(t, err)
	assert.Equal(t, "rpc error: code = Unauthenticated desc = Access token is not provided in authorization metadata", err.Error())
	assert.Nil(t, response)
	assert.Equal(t, 0, mockHandlerCallCounter)
}

func TestUnarityAuthMiddlewareInvalidToken(t *testing.T) {
	token := []string{"invalid-token"}

	ctx := metadata.NewIncomingContext(context.Background(), map[string][]string{"authorization": token})
	info := &grpc.UnaryServerInfo{
		FullMethod: "foo-bar",
	}

	mockHandlerCallCounter := 0
	mockHandler := func(ctx context.Context, req any) (any, error) {
		mockHandlerCallCounter++
		return "foo-bar", nil
	}

	interceptor, err := NewAuthInterceptor(validator, logger, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, interceptor)

	// Should call handler
	response, err := interceptor.UnaryAuthMiddleware(ctx, req, info, mockHandler)

	assert.NotNil(t, err)
	assert.Equal(t, "rpc error: code = Unauthenticated desc = Invalid token. It might be malformed or expired.", err.Error())
	assert.Nil(t, response)
	assert.Equal(t, 0, mockHandlerCallCounter)
}

func TestUnarityAuthMiddlewareValidToken(t *testing.T) {
	token := []string{"valid-token"}

	ctx := metadata.NewIncomingContext(context.Background(), map[string][]string{"authorization": token})
	info := &grpc.UnaryServerInfo{
		FullMethod: "foo-bar",
	}

	mockHandlerCallCounter := 0
	mockHandler := func(ctx context.Context, req any) (any, error) {
		mockHandlerCallCounter++
		assert.Equal(t, ctx.Value(UserIdKey).(string), "1234")
		return "foo-bar", nil
	}

	interceptor, err := NewAuthInterceptor(validator, logger, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, interceptor)

	// Should call handler
	response, err := interceptor.UnaryAuthMiddleware(ctx, req, info, mockHandler)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, mockHandlerCallCounter)
}
