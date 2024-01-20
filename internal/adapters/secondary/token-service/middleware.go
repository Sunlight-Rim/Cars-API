package tokenService

import (
	"cars/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// CheckTokenMW validates and parses token from context, saves claims to context also.
func CheckTokenMW(secret string) echo.MiddlewareFunc {
	byteSecret := []byte(secret)

	return echojwt.WithConfig(echojwt.Config{
		ContextKey: "claims",
		ParseTokenFunc: func(c echo.Context, token string) (interface{}, error) {
			jwtClaims := make(jwt.MapClaims)

			if _, err := jwt.ParseWithClaims(token, jwtClaims, func(*jwt.Token) (any, error) {
				return byteSecret, nil
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
		},
	})
}
