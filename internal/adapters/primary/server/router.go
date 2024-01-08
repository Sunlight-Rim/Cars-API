package server

import (
	"net/http"

	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
)

func (s *server) setRoutes() {
	// Errors list
	s.echo.GET("/errors", func(ctx echo.Context) error {
		return ctx.JSONBlob(http.StatusOK, errors.ResponseList)
	})

	// Health check
	s.echo.GET("/health", func(ctx echo.Context) error {
		return ctx.JSONBlob(http.StatusOK, []byte("Alive!"))
	})

	api := s.echo.Group("/api")

	apiAuth := api.Group("/auth")

	apiAuth.POST("/signup", s.rest.Signup)
}
