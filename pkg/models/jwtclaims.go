package models

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

// valid validates the claims
func (claims JwtClaims) valid() error {
	var now = time.Now().UTC().Unix()
	if claims.UserID == 0 || claims.Email == "" {
		return errors.New("user_id/user_email is required")
	}
	if !claims.VerifyExpiresAt(now, true) {
		return errors.New("token has expired or expiry time is not set")
	}
	return fmt.Errorf("token is invalid")

}
