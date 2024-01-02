package server

import (
	"net/http"

	"cars/pkg/errors"

	openAPI "github.com/go-openapi/runtime/middleware"
	"github.com/labstack/echo/v4"
)

func (s *server) setRoutes() {
	api := s.echo.Group("/api")

	// Swagger
	api.File("/docs/swagger.json", "api/swagger.json")
	api.GET("/docs", echo.WrapHandler(openAPI.SwaggerUI(openAPI.SwaggerUIOpts{
		Path:    "/api/docs",
		SpecURL: "/api/docs/swagger.json",
		Title:   "Cars API Documentation",
	}, nil)))

	// Errors list
	api.GET("/errors", func(ctx echo.Context) error {
		return ctx.JSONBlob(http.StatusOK, errors.ResponseList)
	})

	// api.POST("/signup", s.rest.signup)
}
