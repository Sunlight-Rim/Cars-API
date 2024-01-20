package rest

import (
	"cars/internal/adapters/primary/rest/request"
	"cars/internal/adapters/primary/rest/response"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
)

// CreateCar adds a new car to user.
func (h *Handlers) CreateCar(c echo.Context) (err error) {
	var (
		req *request.CreateCar
		res *response.CreateCar
	)

	// Send response
	defer func() {
		if errResp := c.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewCreateCar(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call use case
	resUcase, err := h.cars.Create(req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "create car")
	}

	res = response.NewCreateCar(resUcase)
	return nil
}

// GetCars

// Update

// DeleteCar
