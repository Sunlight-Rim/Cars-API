package tokenService

import (
	"cars/internal/domain/auth"
	"cars/pkg/errors"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// CheckTokenMW validates and parses token from the context and put the claims in the context.
func CheckTokenMW(secret string) echo.MiddlewareFunc {
	byteSecret := []byte(secret)

	return echojwt.WithConfig(echojwt.Config{
		ContextKey: "claims",
		ParseTokenFunc: func(c echo.Context, token string) (interface{}, error) {
			var claims auth.Claims

			if _, err := jwt.ParseWithClaims(token, &claims, func(*jwt.Token) (any, error) {
				return byteSecret, nil
			}); err != nil {
				if errors.Is(err, jwt.ErrTokenExpired) {
					return nil, errors.Wrapf(errors.ExpiredToken, "expired token, %v", err)
				}
				return nil, errors.Wrapf(errors.InvalidToken, "invalid token, %v", err)
			}

			return &claims, nil
		},
	})
}
