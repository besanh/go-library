package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

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
