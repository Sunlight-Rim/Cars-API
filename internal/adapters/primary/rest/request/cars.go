package request

import (
	"regexp"

	"cars/internal/domain/cars"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
	easyjson "github.com/mailru/easyjson"
)

var plateRegex = regexp.MustCompile(`^[a-z][0-9]{3}[a-z]{2}$`)

// Create car

type CreateCar struct {
	Token string `json:"-"`
	Plate string `json:"plate"`
	Model string `json:"model"`
	Color string `json:"color"`
}

func NewCreateCar(c echo.Context) (*CreateCar, error) {
	r := CreateCar{Token: c.Request().Header.Get("Authorization")}

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	if err := r.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation")
	}

	return &r, nil
}

func (r *CreateCar) Validate() error {
	if !plateRegex.MatchString(r.Plate) {
		return errors.Wrap(errors.InvalidRequestContent, "plate")
	}

	return nil
}

func (r *CreateCar) ToEntity() *cars.CreateReq {
	return &cars.CreateReq{
		Token: r.Token,
		Plate: r.Plate,
		Model: r.Model,
		Color: r.Color,
	}
}
