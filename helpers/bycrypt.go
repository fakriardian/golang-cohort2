package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	cryptFormat = "$%s$%s"
	times       = 1
	memory      = 64 * 1024
	threads     = 4
)

func generateGCM(secret string) cipher.AEAD {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil
	}

	return gcm
}

func GenerateHashPassword(password string) (hash string, err error) {
	cryptoKeyLen, _ := strconv.ParseUint(os.Getenv("AUTH_CRYPTO_KEY_LENGTH"), 10, 64)
	keyLen := uint32(cryptoKeyLen)

	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", nil
	}

	argonHash := argon2.IDKey([]byte(password), salt, times, memory, threads, keyLen)

	b64Hash := encrypt(argonHash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodedHash := fmt.Sprintf(cryptFormat, b64Salt, b64Hash)

	return encodedHash, nil
}

func encrypt(text []byte) string {
	secret := os.Getenv("AUTH_SECRET")
	gcm := generateGCM(secret)
	nonce := make([]byte, gcm.NonceSize())

	cipherText := gcm.Seal(nonce, nonce, text, nil)

	return base64.StdEncoding.EncodeToString(cipherText)
}

func decrypt(cipherText string) ([]byte, error) {
	secret := os.Getenv("AUTH_SECRET")
	gcm := generateGCM(secret)

	decoded, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	if len(decoded) < gcm.NonceSize() {
		return nil, errors.New("invalid nonce size")
	}

	return gcm.Open(
		nil,
		decoded[:gcm.NonceSize()],
		decoded[gcm.NonceSize():],
		nil,
	)
}

func ComparePassword(password, hash string) (bool, error) {
	parts := strings.Split(hash, "$")
	salt, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	hash = parts[2]

	decryptHash, err := decrypt(hash)
	if err != nil {
		return false, err
	}

	var keyLen = uint32(len(decryptHash))

	comparisionHash := argon2.IDKey([]byte(password), salt, times, memory, threads, keyLen)

	return subtle.ConstantTimeCompare(comparisionHash, decryptHash) == 1, nil
}
