package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Udehlee/Task-Management/pkg/models"
	"github.com/golang-jwt/jwt/v4"
)

// avoid context key collisions
type contextKey string

const claimsKey = contextKey("claims")

// AuthMiddleware extracts and validates the token, then stores the claims in the context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from request header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := ValidateToken(tokenString, "your-secret-key")
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Store claims in request context
		ctx := context.WithValue(r.Context(), claimsKey, claims)
		r = r.WithContext(ctx)

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

// ValidateToken validates the JWT token and returns the claims
func ValidateToken(tokenString string, TokenKey string) (*models.JwtClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &models.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(TokenKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, err := GetClaims(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// GetClaims extracts the claims from the token
func GetClaims(token *jwt.Token) (*models.JwtClaims, error) {

	claims, ok := token.Claims.(*models.JwtClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims type")
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}
	return claims, nil
}
