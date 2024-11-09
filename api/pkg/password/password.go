package password

/// Password implementation from snyk.io using Argon2id
// Argon2id is the OWASP recommended algorithm for hashing password at the moment of writing this
/// https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html
import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/argon2"
)

type HashSalt struct {
	Hash string
	Salt string
}

// Argon2id params
type Password struct {
	// time represents the number of passed over the specified memory.
	time uint32
	// cpu memory to be used.
	memory uint32
	// threads for parallelism aspect of the algorithm.
	threads uint8
	// keyLen of the generate hash key.
	keyLen uint32
	// saltLen the length of the salt used.
	saltLen uint32
}

func New() *Password {
	// Set default values for Argon2id
	return &Password{
		time:    1,
		saltLen: 32,
		memory:  64 * 1024,
		threads: 32,
		keyLen:  256,
	}
}

func randomSecret(length uint32) ([]byte, error) {
	secret := make([]byte, length)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

// GenerateHash using the password and provided salt.
// If not salt value provided fallback to random value
// generated of a given length.
func (a *Password) GenerateHash(password, salt []byte) (*HashSalt, error) {
	var err error
	// If salt is not provided generate a salt of
	// the configured salt length.
	if len(salt) == 0 {
		salt, err = randomSecret(a.saltLen)
	}

	if err != nil {
		return nil, err
	}

	// Generate hash
	hash := argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen)

	// Base64 encode the salt and hashed password to persist it
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return the generated hash and salt used for storage.
	return &HashSalt{Hash: b64Hash, Salt: b64Salt}, nil
}

// Compare stored hash (with salt) with provided password
func (a *Password) Compare(existingHash, existingSalt, password string) error {
	existingSaltInBytes, err := base64.RawStdEncoding.DecodeString(existingSalt)
	if err != nil {
		return err
	}

	// Convert given password to hash, use stored salt
	hashedPassword, err := a.GenerateHash([]byte(password), existingSaltInBytes)

	if err != nil {
		return err
	}

	if hashedPassword.Hash != existingHash {
		return errors.New("hash doesn't match")
	}
	return nil
}
