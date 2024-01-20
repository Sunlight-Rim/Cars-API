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

func (h *Handlers) Register(api *echo.Group) {
	// Auth

	apiAuth := api.Group("/auth")
	/*
		swagger:route POST /api/auth/signup Auth SignupRequest

		Register a new user.
		Password must contain at least one special character in a range [.,\(\);:\\\/\[\]\{\}@$!%*#?&=].

			schemes: http
			responses:
				200: SignupResponse
				default: ErrorResponse
	*/
	apiAuth.POST("/signup", h.Signup)
	/*
		swagger:route POST /api/auth/signin Auth SigninRequest

		Sign in to user account.

			schemes: http
			responses:
				200: SigninResponse
				default: ErrorResponse
	*/
	apiAuth.POST("/signin", h.Signin)
	/*
		swagger:route POST /api/auth/refresh Auth RefreshRequest

		Create new token from expired, also revoke expired.

			schemes: http
			responses:
				200: RefreshResponse
				default: ErrorResponse
	*/
	apiAuth.POST("/refresh", h.Refresh)
	/*
		swagger:route POST /api/auth/signout Auth SignoutRequest

		Revoke refresh token.

			schemes: http
			responses:
				200: SignoutResponse
				default: ErrorResponse
	*/
	apiAuth.POST("/signout", h.Signout)
	// apiAuth.POST("/signout-all", h.SignoutAll)

	// User

	apiUser := api.Group("/user")
	/*
		swagger:route GET /api/user User null

		Sign in to account.

			schemes: http
			security:
				accessToken: []
			responses:
				200: GetMeResponse
				default: ErrorResponse
	*/
	apiUser.GET("", h.GetMe)
	// apiUser.PUT("/info", h.UpdateUserInfo)
	// apiUser.PUT("/password", h.UpdatePassword)
	// apiUser.DELETE("", h.DeleteUser)

	// Cars

	apiCars := api.Group("/cars")
	/*
		swagger:route POST /api/cars Cars CreateCarRequest

		Create car on user account.

			schemes: http
			security:
				accessToken: []
			responses:
				200: CreateCarResponse
				default: ErrorResponse
	*/
	apiCars.POST("", h.CreateCar)
	// apiCars.GET("", h.GetCars)
	// apiCars.PUT("/color", h.UpdateCarColor)
	// apiCars.DELETE("", h.DeleteCar)
}
