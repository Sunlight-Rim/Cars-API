/*
Package app Cars API

API with CRUD operations on Users and their Cars.

		version: 0.1
		host: localhost:1337
		schemes: http
	    consumes:
	    - application/json

	    produces:
	    - application/json

	    securityDefinitions:
	    accessToken:
	        type: apiKey
	        name: Authorization
	        in: cookie
	        description: JWT authorization token stored in a cookie.

swagger:meta
*/
package server

import (
	"net/http"

	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
)

// Register services handlers
func (s *server) Register() {
	// REST handlers

	s.rest.Register(s.echo.Group("/api"))

	// GraphQL handlers

	s.graphql.Register(s.echo.Group("/graphql"))

	// General handlers

	/*
		swagger:route GET /errors Errors none

		List of API errors.

			schemes: https
			responses:
				200: ErrorsListResponse
				default: ErrorResponse
	*/
	s.echo.GET("/errors", func(ctx echo.Context) error {
		return ctx.JSONBlob(http.StatusOK, errors.ResponseList)
	})
	/*
		swagger:route GET /health Health none

		Health check.

			schemes: https
			responses:
				200: body:string
				default: ErrorResponse
	*/
	s.echo.GET("/health", func(ctx echo.Context) error {
		return ctx.JSONBlob(http.StatusOK, []byte("Alive!"))
	})
}
