package utils

import (
	"strconv"
	"time"

	"github.com/Udehlee/Task-Management/pkg/models"
	"github.com/golang-jwt/jwt"
)

var (
	jwtKey = []byte("YOUR_JWT_SECRET_KEY")
)

func GenerateToken(user models.User) (string, error) {
	// Create the JWT claims, which include the user ID and expiry time
	claims := models.JwtClaims{
		UserID:    user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		StandardClaims: jwt.StandardClaims{
			Subject:   strconv.Itoa(user.UserID),
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(), // Token expires in 10 hours
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create a new JWT token object using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "Jwt creation failed", err
	}

	return tokenString, nil
}
