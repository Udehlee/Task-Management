package utils

import (
	"strconv"
	"testing"
	"time"

	"github.com/Udehlee/Task-Management/pkg/models"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	//mock user
	user := models.User{
		UserID:    1,
		FirstName: "Nma",
		LastName:  "James",
		Email:     "nmaj@gmail.com",
	}

	//generate token
	tokenString, err := GenerateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Parse the token
	//jwtKey already declared in authToken
	token, err := jwt.ParseWithClaims(tokenString, &models.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	assert.NoError(t, err)
	assert.NotNil(t, token)

	// Verify the token claims
	claims, ok := token.Claims.(*models.JwtClaims)
	assert.True(t, ok)
	assert.True(t, token.Valid)
	assert.Equal(t, user.UserID, claims.UserID)
	assert.Equal(t, user.FirstName, claims.FirstName)
	assert.Equal(t, user.LastName, claims.LastName)
	assert.Equal(t, strconv.Itoa(user.UserID), claims.Subject)
	assert.WithinDuration(t, time.Now().Add(time.Hour*10), time.Unix(claims.ExpiresAt, 0), time.Minute)
	assert.WithinDuration(t, time.Now(), time.Unix(claims.IssuedAt, 0), time.Minute)
}
