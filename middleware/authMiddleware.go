package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Udehlee/Task-Management/pkg/models"
	"github.com/Udehlee/Task-Management/utils"
	"github.com/golang-jwt/jwt/v4"
)

// avoid context key collisions
type contextKey string

const userIDKey = contextKey("UserID")

// AuthMiddleware extracts and validates the token, then stores the claims in the context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from request header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.UnsucessfulRequest(w, "Unauthorized", "Authorisation token is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := ValidateToken(tokenString, "YOUR_JWT_SECRET_KEY")
		if err != nil {
			utils.UnsucessfulRequest(w, "Unauthorized", "Invalid token"+err.Error(), http.StatusUnauthorized)
			return
		}

		// Extract user ID from claims
		userID := claims.UserID
		if userID == 0 {
			log.Fatal("user not found")
		}

		// Store claims in request context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
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
		return []byte("YOUR_JWT_SECRET_KEY"), nil
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
