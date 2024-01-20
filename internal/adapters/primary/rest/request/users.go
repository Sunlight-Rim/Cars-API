package request

import (
	"cars/internal/domain/users"

	"github.com/labstack/echo/v4"
)

// Get me

type GetMe struct {
	Token string
}

func NewGetMe(c echo.Context) (*GetMe, error) {
	return &GetMe{Token: c.Request().Header.Get("Authorization")}, nil
}

func (r *GetMe) ToEntity() *users.GetReq {
	return &users.GetReq{
		Token: r.Token,
	}
}
