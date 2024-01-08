/*
Package app Cars API

API with CRUD-operations on Users and his Cars.

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

func (s *server) setRoutes() {
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
				200: HealthResponse
				default: ErrorResponse
	*/
	s.echo.GET("/health", func(ctx echo.Context) error {
		return ctx.JSONBlob(http.StatusOK, []byte("Alive!"))
	})

	api := s.echo.Group("/api")

	apiAuth := api.Group("/auth")
	/*
		swagger:route POST /api/auth/signup Auth SignupRequest

		Register a new user account.
		Password must contain at least one special character in a range [.,\(\);:\\\/\[\]\{\}@$!%*#?&=].

			schemes: https
			responses:
				200: SignupResponse
				default: ErrorResponse
	*/
	apiAuth.POST("/signup", s.rest.Signup)
	// apiAuth.POST("/signin", s.rest.Signin)
	// apiAuth.POST("/signout", s.rest.Signout)
	// apiAuth.POST("/refresh", s.rest.Refresh)

	// apiUser := api.Group("/user")
	// apiUser.GET("", s.rest.GetUser)
	// apiUser.PUT("/info", s.rest.UpdateUserInfo)
	// apiUser.PUT("/password", s.rest.UpdatePassword)
	// apiUser.DELETE("", s.rest.DeleteUser)

	// apiCars := api.Group("/cars")
	// apiCars.POST("", s.rest.CreateCar)
	// apiCars.GET("", s.rest.GetCar)
	// apiCars.PUT("/color", s.rest.UpdateCarColor)
	// apiCars.DELETE("", s.rest.DeleteCar)
}
