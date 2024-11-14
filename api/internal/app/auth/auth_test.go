package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var secret = "test-secret-key"

func TestNewAuth(t *testing.T) {
	auth, err := New(secret)
	assert.Nil(t, err)
	assert.NotNil(t, auth)
	assert.NotNil(t, auth.Password)
}

func TestTokenValid(t *testing.T) {
	auth, _ := New(secret)
	tokens, err := auth.IssueToken(1)

	assert.Nil(t, err)
	assert.NotNil(t, tokens.AccessToken)
	assert.NotNil(t, tokens.RefreshToken)

	userId, err := auth.ValidateToken(tokens.AccessToken)
	assert.Nil(t, err)
	assert.Equal(t, userId, "1")
}

func TestTokenValidationError(t *testing.T) {
	auth, _ := New(secret)
	tokens, err := auth.IssueToken(1)

	assert.Nil(t, err)
	assert.NotNil(t, tokens.AccessToken)
	assert.NotNil(t, tokens.RefreshToken)

	userId, err := auth.ValidateToken(tokens.AccessToken + "foo")
	assert.Empty(t, userId)
	assert.Error(t, err)
}

func TestGenerateRandomString(t *testing.T) {
	str, err := GenerateRandomString(32)
	assert.Nil(t, err)
	assert.NotNil(t, str)
	assert.Len(t, str, 44)
}
