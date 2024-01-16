package tokenService

import (
	"cars/internal/domain/auth"
	"cars/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
)

type parser struct {
	secret string
}

func NewParser(secret string) *parser {
	return &parser{
		secret: secret,
	}
}

func (s *parser) Parse(token string) (auth.Claims, error) {
	claims := jwt.MapClaims(auth.Claims{})

	if _, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (any, error) {
		return s.secret, nil
	}); err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.Wrapf(errors.ExpiredToken, "expired token, %v", err)
		}
		return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
	}

	return claims, nil
}
