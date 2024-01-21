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

	checkTokenMW echo.MiddlewareFunc
}

func New(auth auth.IUsecase, users users.IUsecase, cars cars.IUsecase, checkTokenMW echo.MiddlewareFunc) *Handlers {
	return &Handlers{
		auth:  auth,
		users: users,
		cars:  cars,

		checkTokenMW: checkTokenMW,
	}
}

func (h *Handlers) Register(api *echo.Group) {

	// Auth

	apiAuth := api.Group("/auth")

	/*
		swagger:route POST /api/auth/signup Auth SignupRequest

		Register a new user.
		Password must be longer than 6 characters and contain at least one special character.

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
	/*
		swagger:route POST /api/auth/signout-all Auth SignoutAllRequest

		Revoke all user refresh tokens.

			schemes: http
			responses:
				200: SignoutAllResponse
				default: ErrorResponse
	*/
	apiAuth.POST("/signout-all", h.SignoutAll)

	// User

	apiUser := api.Group("/user", h.checkTokenMW)

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
	// apiUser.PUT("", h.UpdateUser)
	// apiUser.PUT("/password", h.UpdatePassword)
	// apiUser.DELETE("", h.DeleteUser)

	// Cars

	apiCars := api.Group("/cars", h.checkTokenMW)

	/*
		swagger:route POST /api/cars Cars CreateCarRequest

		Create car in user account.
		Plate must be in format of three latin letters, then three digits. For example "aaa111".

			schemes: http
			security:
				accessToken: []
			responses:
				200: CreateCarResponse
				default: ErrorResponse
	*/
	apiCars.POST("", h.CreateCar)
	/*
		swagger:route GET /api/cars Cars null

		Get all user cars.

			schemes: http
			security:
				accessToken: []
			responses:
				200: GetCarsResponse
				default: ErrorResponse
	*/
	apiCars.GET("", h.GetCars)
	/*
		swagger:route PUT /api/cars Cars UpdateCarRequest

		Get all user cars.

			schemes: http
			security:
				accessToken: []
			responses:
				200: UpdateCarResponse
				default: ErrorResponse
	*/
	apiCars.PUT("", h.UpdateCar)
	/*
		swagger:route DELETE /api/cars Cars DeleteCarRequest

		Get all user cars.

			schemes: http
			security:
				accessToken: []
			responses:
				200: DeleteCarResponse
				default: ErrorResponse
	*/
	apiCars.DELETE("", h.DeleteCar)
}
