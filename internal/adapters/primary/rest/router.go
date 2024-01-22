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
	apiAuth.POST("/signup", h.signup)
	/*
		swagger:route POST /api/auth/signin Auth SigninRequest

		Sign in to user account.

			schemes: http
			responses:
				200: SigninResponse
				default: ErrorResponse
	*/
	apiAuth.POST("/signin", h.signin)
	/*
		swagger:route POST /api/auth/refresh Auth RefreshRequest

		Create new token from expired and revoke expired.

			schemes: http
			responses:
				200: RefreshResponse
				default: ErrorResponse
	*/
	apiAuth.POST("/refresh", h.refresh)
	/*
		swagger:route POST /api/auth/sessions Auth SessionsRequest

		List all user active sessions (refresh tokens).

			schemes: http
			responses:
				200: SessionsResponse
				default: ErrorResponse
	*/
	apiAuth.POST("/sessions", h.sessions)
	/*
		swagger:route POST /api/auth/signout Auth SignoutRequest

		Revoke refresh token.

			schemes: http
			responses:
				200: SignoutResponse
				default: ErrorResponse
	*/
	apiAuth.POST("/signout", h.signout)
	/*
		swagger:route POST /api/auth/signout-all Auth SignoutAllRequest

		Revoke all user refresh tokens.

			schemes: http
			responses:
				200: SignoutAllResponse
				default: ErrorResponse
	*/
	apiAuth.POST("/signout-all", h.signoutAll)

	// User

	apiUser := api.Group("/user", h.checkTokenMW)

	/*
		swagger:route GET /api/user User GetMeNull

		Get user account information.

			schemes: http
			security:
				accessToken: []
			responses:
				200: GetMeResponse
				default: ErrorResponse
	*/
	apiUser.GET("", h.getMe)
	/*
		swagger:route PUT /api/user User UpdateInfoRequest

		Update user general information.

			schemes: http
			security:
				accessToken: []
			responses:
				200: UpdateInfoResponse
				default: ErrorResponse
	*/
	apiUser.PUT("", h.updateInfo)
	/*
		swagger:route PUT /api/user/password User UpdatePasswordRequest

		Change user password.

			schemes: http
			security:
				accessToken: []
			responses:
				200: UpdatePasswordResponse
				default: ErrorResponse
	*/
	apiUser.PUT("/password", h.updatePassword)
	/*
		swagger:route DELETE /api/user User DeleteUserNull

		Delete user account.

			schemes: http
			security:
				accessToken: []
			responses:
				200: DeleteUserResponse
				default: ErrorResponse
	*/
	apiUser.DELETE("", h.deleteUser)

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
	apiCars.POST("", h.createCar)
	/*
		swagger:route GET /api/cars Cars GetCarsNull

		Get all user cars.

			schemes: http
			security:
				accessToken: []
			responses:
				200: GetCarsResponse
				default: ErrorResponse
	*/
	apiCars.GET("", h.getCars)
	/*
		swagger:route PUT /api/cars Cars UpdateCarRequest

		Update car data.

			schemes: http
			security:
				accessToken: []
			responses:
				200: UpdateCarResponse
				default: ErrorResponse
	*/
	apiCars.PUT("", h.updateCar)
	/*
		swagger:route DELETE /api/cars Cars DeleteCarRequest

		Delete user car.

			schemes: http
			security:
				accessToken: []
			responses:
				200: DeleteCarResponse
				default: ErrorResponse
	*/
	apiCars.DELETE("", h.deleteCar)
}
