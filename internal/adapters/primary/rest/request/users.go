package request

import (
	"cars/internal/domain/auth"
	"cars/internal/domain/users"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
	easyjson "github.com/mailru/easyjson"
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

// Update

type UpdateInfo struct {
	UserID   uint64 `json:"-"`
	Username string `json:"username"`
	Phone    uint64 `json:"phone"`
}

func NewUpdateInfo(c echo.Context) (*UpdateInfo, error) {
	r := UpdateInfo{UserID: c.Get("claims").(*auth.Claims).UserID}

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}

func (r *UpdateInfo) ToEntity() *users.UpdateInfoReq {
	return &users.UpdateInfoReq{
		UserID:   r.UserID,
		Username: r.Username,
		Phone:    r.Phone,
	}
}

// Update password

type UpdatePassword struct {
	UserID      uint64 `json:"-"`
	NewPassword string `json:"new_password"`
}

func NewUpdatePassword(c echo.Context) (*UpdatePassword, error) {
	r := UpdatePassword{UserID: c.Get("claims").(*auth.Claims).UserID}

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}

func (r *UpdatePassword) ToEntity() *users.UpdatePasswordReq {
	return &users.UpdatePasswordReq{
		UserID:      r.UserID,
		NewPassword: r.NewPassword,
	}
}

// Delete

type DeleteUser struct {
	UserID uint64 `json:"-"`
}

func NewDeleteUser(c echo.Context) (*DeleteUser, error) {
	return &DeleteUser{UserID: c.Get("claims").(*auth.Claims).UserID}, nil
}

func (r *DeleteUser) ToEntity() *users.DeleteReq {
	return &users.DeleteReq{
		UserID: r.UserID,
	}
}
