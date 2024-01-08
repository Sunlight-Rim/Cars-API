package request

import (
	authEntity "cars/internal/entities/auth"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
)

// Signup

func NewSignup(c echo.Context) (*Signup, error) {
	var r Signup

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	if err := r.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation")
	}

	return &r, nil
}

func (r *Signup) Validate() error {
	return nil
}

func (r *Signup) ToEntity() *authEntity.UsecaseReqSignup {
	return &authEntity.UsecaseReqSignup{
		Username: r.Username,
		Email:    r.Email,
		Address:  r.Address,
	}
}
