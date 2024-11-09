package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ossipesonen/go-traffic-lights/pkg/password"
)

type IssuedTokens struct {
	AccessToken  string
	RefreshToken string
}

type Auth struct {
	secret []byte
	// Allow access to use Password utilities
	Password *password.Password
}

var ErrInvalidToken = errors.New("invalid token")

func New(secret string) (*Auth, error) {
	if secret == "" {
		return nil, errors.New("cannot have an empty secret")
	}

	return &Auth{secret: []byte(secret), Password: password.New()}, nil
}

// IssueToken will issue a JWT token with the provided userID as the subject. The token will expire after 60 minutes.
func (s *Auth) IssueToken(userID int) (*IssuedTokens, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(userID),                    // RFC defines "sub" should be string
		"iss": "traffic-lights",                        // Todo: do not hard code this
		"exp": time.Now().Add(time.Minute * 60).Unix(), // 60 min
	}, nil)

	signedAccessToken, err := accessToken.SignedString(s.secret)
	if err != nil {
		return nil, fmt.Errorf("failed to sign access token: %w", err)
	}

	// Refresh token can be opaque as we store it
	refreshToken, err := GenerateRandomString(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate a refresh token: %w", err)
	}

	return &IssuedTokens{AccessToken: signedAccessToken, RefreshToken: refreshToken}, nil
}

// ValidateToken will validate the provide JWT against the secret key. It'll then check if the token has expired, and then return the user ID set as the token subject.
func (s *Auth) ValidateToken(token string) (string, error) {
	// validate token for the correct secret key and signing method.
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secret, nil
	})

	if err != nil {
		return "", errors.Join(ErrInvalidToken, err)
	}

	// read claims from payload and extract the user ID.
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		id, ok := claims["sub"].(string)
		if !ok {
			return "", fmt.Errorf("%w: failed to extract id from claims", ErrInvalidToken)
		}

		return id, nil
	}

	return "", ErrInvalidToken
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
