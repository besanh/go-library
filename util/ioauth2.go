package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

var KEY_AES_256 string = "0123456789ABCDEF0123456789ABCDEF"

// GenerateCodeVerifier creates a random string as the code verifier.
func (i *Util) GenerateCodeVerifier() (string, error) {
	verifier := make([]byte, 32)
	_, err := rand.Read(verifier)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(verifier), nil
}

// GenerateCodeChallenge generates a code challenge using SHA256.
func (i *Util) GenerateCodeChallenge(verifier string) string {
	hash := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

// Encrypt encrypts the plaintext using AES-GCM with the provided key.
// It returns the ciphertext encoded in base64.
func (i *Util) Encrypt(plaintext string) (string, error) {
	// Create a new cipher block with the provided key.
	block, err := aes.NewCipher([]byte(KEY_AES_256))
	if err != nil {
		return "", err
	}

	// Create an AES-GCM instance from the block cipher.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Generate a random nonce of the required size.
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the plaintext. The nonce is prepended to the ciphertext.
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)

	// Return the encrypted text as a base64 encoded string.
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the base64 encoded ciphertext using AES-GCM with the provided key.
// It returns the original plaintext.
func (i *Util) Decrypt(encryptedText string) (string, error) {
	// Decode the base64 encoded ciphertext.
	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	// Create a new cipher block using the provided key.
	block, err := aes.NewCipher([]byte((KEY_AES_256)))
	if err != nil {
		return "", err
	}

	// Create an AES-GCM instance from the block cipher.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	// Extract the nonce and the actual ciphertext.
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	// Decrypt the ciphertext.
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
