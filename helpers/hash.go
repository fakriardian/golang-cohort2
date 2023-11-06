package helpers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
	"strconv"
	"strings"
)

const encSuffix = "ASD:"

type Encryption struct {
	key       []byte
	algorithm string
	iv        []byte
}

func NewEncryption() *Encryption {
	secretKey := os.Getenv("AUTH_SECRET")
	keyLength, _ := strconv.Atoi(os.Getenv("AUTH_CRYPTO_KEY_LENGTH"))
	algorithm := os.Getenv("AUTH_CRYPTO_ALGORITHM")
	ivLength, _ := strconv.Atoi(os.Getenv("AUTH_CRYPTO_IV_LENGTH"))

	key := []byte(secretKey)[:keyLength]
	iv := []byte(secretKey)[:ivLength]

	return &Encryption{
		key:       key,
		algorithm: algorithm,
		iv:        iv,
	}
}

func (e *Encryption) Encrypt(value string) string {
	cipherBlock, err := aes.NewCipher(e.key)
	if err != nil {
		return ""
	}

	paddedValue := padValue([]byte(value), aes.BlockSize)
	cipherText := make([]byte, len(paddedValue))
	encrypter := cipher.NewCBCEncrypter(cipherBlock, e.iv)
	encrypter.CryptBlocks(cipherText, paddedValue)

	result := encSuffix + base64.StdEncoding.EncodeToString(cipherText)
	return result
}

func (e *Encryption) Decrypt(token string) string {
	if !strings.HasPrefix(token, encSuffix) {
		return token
	}

	cipherText, err := base64.StdEncoding.DecodeString(token[len(encSuffix):])
	if err != nil {
		return ""
	}

	cipherBlock, err := aes.NewCipher(e.key)
	if err != nil {
		return ""
	}

	decrypter := cipher.NewCBCDecrypter(cipherBlock, e.iv)
	decrypter.CryptBlocks(cipherText, cipherText)

	unpaddedValue := unpadValue(cipherText)
	result := string(unpaddedValue)
	return result
}

func padValue(value []byte, blockSize int) []byte {
	padding := blockSize - (len(value) % blockSize)
	paddedValue := append(value, bytes.Repeat([]byte{byte(padding)}, padding)...)
	return paddedValue
}

func unpadValue(paddedValue []byte) []byte {
	padding := int(paddedValue[len(paddedValue)-1])
	if padding <= 0 || padding > aes.BlockSize {
		return paddedValue
	}
	unpaddedValue := paddedValue[:len(paddedValue)-padding]
	return unpaddedValue
}
