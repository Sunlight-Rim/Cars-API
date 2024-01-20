package request

import (
	"cars/internal/domain/auth"
	"cars/internal/domain/users"

	"github.com/labstack/echo/v4"
)

// Get me

type GetMe struct {
	UserID uint64 `json:"-"`
}

func NewGetMe(c echo.Context) (*GetMe, error) {
	return &GetMe{UserID: c.Get("claims").(*auth.Claims).UserID}, nil
}

func (r *GetMe) ToEntity() *users.GetReq {
	return &users.GetReq{
		UserID: r.UserID,
	}
}
