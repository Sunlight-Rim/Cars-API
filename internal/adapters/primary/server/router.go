package server

import (
	"net/http"

	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
)

func (s *server) setRoutes() {
	api := s.echo.Group("/api")

	// Errors list
	api.GET("/errors", func(ctx echo.Context) error {
		return ctx.JSONBlob(http.StatusOK, errors.ResponseList)
	})

	// Health check
	api.GET("/health", func(ctx echo.Context) error {
		return ctx.JSONBlob(http.StatusOK, []byte("Alive!"))
	})

	// api.POST("/signup", s.rest.signup)
}
