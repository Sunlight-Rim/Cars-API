package rest

import (
	"cars/internal/domain/auth"
	"cars/internal/domain/cars"
	"cars/internal/domain/users"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	auth  auth.IUsecase
	users users.IUsecase
	cars  cars.IUsecase
}

func New(auth auth.IUsecase, users users.IUsecase, cars cars.IUsecase) *Handlers {
	return &Handlers{
		auth:  auth,
		users: users,
		cars:  cars,
	}
}

func (h Handlers) Register(api *echo.Group) {
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
	apiAuth.POST("/signup", h.Signup)
	/*
		swagger:route POST /api/auth/signin Auth SigninRequest

		Sign in to account.

			schemes: https
			responses:
				200: SigninResponse
				default: ErrorResponse
	*/
	apiAuth.POST("/signin", h.Signin)
	// apiAuth.POST("/refresh", s.rest.Refresh)
	// apiAuth.POST("/signout", s.rest.Signout)
	// apiAuth.POST("/signout-all", s.rest.SignoutAll)

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
