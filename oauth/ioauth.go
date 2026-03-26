package oauth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/besanh/go-library/oauth/context"
	"github.com/golang-jwt/jwt/v5"
)

type IOauth interface {
	ParseAndVerify(tokenString string) (*jwt.Token, error)
	Close()
	StandardMiddleware(next http.Handler) http.Handler
}

// ParseAndVerify takes a raw token string and verifies it against the JWKS
func (a *Authenticator) ParseAndVerify(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Enforce Asymmetric RSA algorithm
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the dynamically fetched public key
		return a.jwks.Keyfunc(token)
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

// Close cleans up background refresh goroutines
func (a *Authenticator) Close() {
	if a.jwks != nil {
		a.jwks.EndBackground()
	}
}

// StandardMiddleware returns a standard net/http middleware
func (a *Authenticator) StandardMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Extract bearer token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "authorization header missing", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// 2. Verify token
		token, err := a.ParseAndVerify(tokenString)
		if err != nil {
			http.Error(w, "unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// 3. Extract Subject (User ID) from Claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "invalid token claims", http.StatusUnauthorized)
			return
		}

		// 4. Extract User ID from claims
		userID, ok := claims["sub"].(string)
		if !ok || userID == "" {
			http.Error(w, "user ID missing in token", http.StatusUnauthorized)
			return
		}

		// 5. Inject User ID into context
		ctx := context.InjectUserID(r.Context(), userID)

		// 6. Call next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
