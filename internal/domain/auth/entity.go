package auth

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID uint64

	jwt.RegisteredClaims
}
