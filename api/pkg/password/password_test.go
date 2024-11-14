package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	password := New()
	assert.NotNil(t, password)
}

func TestGenerateHash(t *testing.T) {
	password := New()
	assert.NotNil(t, password)

	hashSalt, err := password.GenerateHash([]byte("test-password"), []byte("test-salt"))

	assert.Nil(t, err)
	assert.NotNil(t, hashSalt.Hash)
	assert.NotNil(t, hashSalt.Salt)

	assert.Len(t, hashSalt.Hash, 342)
	assert.Len(t, hashSalt.Salt, 12)
}

func TestGenerateHashNoSalt(t *testing.T) {
	password := New()
	assert.NotNil(t, password)

	var salt []byte
	hashSalt, err := password.GenerateHash([]byte("test-password"), salt)

	assert.Nil(t, err)
	assert.NotNil(t, hashSalt.Hash)
	assert.NotNil(t, hashSalt.Salt)

	assert.Len(t, hashSalt.Hash, 342)
	assert.Len(t, hashSalt.Salt, 43)
}

func TestCompareSuccess(t *testing.T) {
	password := New()
	assert.NotNil(t, password)

	hashSalt, _ := password.GenerateHash([]byte("test-password"), []byte("test-salt"))
	err := password.Compare(hashSalt.Hash, hashSalt.Salt, "test-password")

	assert.Nil(t, err)
}

func TestCompareFailure(t *testing.T) {
	password := New()
	assert.NotNil(t, password)

	hashSalt, _ := password.GenerateHash([]byte("test-password"), []byte("test-salt"))
	err := password.Compare(hashSalt.Hash, hashSalt.Salt, "test-password")

	assert.Nil(t, err)
}
