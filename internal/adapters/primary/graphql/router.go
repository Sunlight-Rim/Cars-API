package graphql

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

func (h *Handlers) Register(api *echo.Group) {}
