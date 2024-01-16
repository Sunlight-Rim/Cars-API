package tokenService

import (
	"cars/internal/domain/auth"
	"cars/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
)

type parser struct {
	secret []byte
}

func NewParser(secret string) *parser {
	return &parser{
		secret: []byte(secret),
	}
}

// Parse token.
func (s *parser) Parse(token string) (*auth.Claims, error) {
	jwtClaims := make(jwt.MapClaims)

	if _, err := jwt.ParseWithClaims(token, jwtClaims, func(*jwt.Token) (any, error) {
		return s.secret, nil
	}); err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.Wrapf(errors.ExpiredToken, "expired token, %v", err)
		}
		return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
	}

	claims, err := parse(jwtClaims)
	if err != nil {
		return nil, errors.Wrap(err, "parse claims")
	}

	return claims, nil
}
