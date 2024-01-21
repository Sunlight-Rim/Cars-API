package request

import (
	"regexp"

	"cars/internal/domain/auth"
	"cars/internal/domain/cars"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
	easyjson "github.com/mailru/easyjson"
)

var plateRegex = regexp.MustCompile(`^[a-z]{3}[0-9]{3}$`)

// Create

type CreateCar struct {
	UserID uint64 `json:"-"`
	Plate  string `json:"plate"`
	Model  string `json:"model"`
	Color  string `json:"color"`
}

func NewCreateCar(c echo.Context) (*CreateCar, error) {
	r := CreateCar{UserID: c.Get("claims").(*auth.Claims).UserID}

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
		UserID: r.UserID,
		Plate:  r.Plate,
		Model:  r.Model,
		Color:  r.Color,
	}
}

// Get

type GetCars struct {
	UserID uint64 `json:"-"`
}

func NewGetCars(c echo.Context) *GetCars {
	return &GetCars{UserID: c.Get("claims").(*auth.Claims).UserID}
}

func (r *GetCars) ToEntity() *cars.GetReq {
	return &cars.GetReq{
		UserID: r.UserID,
	}
}

// Update

type UpdateCar struct {
	UserID uint64 `json:"-"`
	CarID  uint64 `json:"car_id"`
	Model  string `json:"model"`
	Color  string `json:"color"`
}

func NewUpdateCar(c echo.Context) (*UpdateCar, error) {
	r := UpdateCar{UserID: c.Get("claims").(*auth.Claims).UserID}

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}

func (r *UpdateCar) ToEntity() *cars.UpdateReq {
	return &cars.UpdateReq{
		UserID: r.UserID,
		CarID:  r.CarID,
		Model:  r.Model,
		Color:  r.Color,
	}
}

// Delete

type DeleteCar struct {
	UserID uint64 `json:"-"`
	CarID  uint64 `json:"car_id"`
}

func NewDeleteCar(c echo.Context) (*DeleteCar, error) {
	r := DeleteCar{UserID: c.Get("claims").(*auth.Claims).UserID}

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}

func (r *DeleteCar) ToEntity() *cars.DeleteReq {
	return &cars.DeleteReq{
		UserID: r.UserID,
		CarID:  r.CarID,
	}
}
