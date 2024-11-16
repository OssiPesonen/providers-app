package interceptor

import (
	"context"
	"errors"
	"log"
	"slices"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type key string

const UserIdKey key = "user_id"

type (
	// Validator defines an interface for token validation. This is satisfied by the AuthService.
	Validator interface {
		ValidateToken(token string) (string, error)
	}

	AuthInterceptor struct {
		validator          Validator
		logger             *log.Logger
		whitelistedMethods []string
	}
)

func NewAuthInterceptor(validator Validator, logger *log.Logger, whitelistedRoutes []string) (*AuthInterceptor, error) {
	if validator == nil {
		return nil, errors.New("validator cannot be nil")
	}

	return &AuthInterceptor{validator: validator, logger: logger, whitelistedMethods: whitelistedRoutes}, nil
}

func (a *AuthInterceptor) UnaryAuthMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// Grab metadata for authorization key
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Metadata has not been provided")
	}

	// Skip defined routes
	if len(a.whitelistedMethods) > 0 {
		// Full method is typically /proto.{service}/{method}
		// so just compare with method name
		method := info.FullMethod[strings.LastIndex(info.FullMethod, "/")+1:]

		if slices.Contains(a.whitelistedMethods, method) {
			return handler(ctx, req)
		}
	}

	// extract token from authorization key
	token := md["authorization"]
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Access token is not provided in authorization metadata")
	}

	t, found := strings.CutPrefix(token[0], "Bearer ")
	if !found {
		return nil, status.Error(codes.Unauthenticated, "Bearer token is not provide in the authorization metadata")
	}

	// validate token and retrieve the userID
	userID, err := a.validator.ValidateToken(t)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid token. It might be malformed or expired.")
	}

	// add our user ID to the context, so we can use it in our RPC handler
	ctx = context.WithValue(ctx, UserIdKey, userID)

	// call our handler
	return handler(ctx, req)
}
