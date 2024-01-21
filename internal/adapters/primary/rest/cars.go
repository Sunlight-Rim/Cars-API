package rest

import (
	"cars/internal/adapters/primary/rest/request"
	"cars/internal/adapters/primary/rest/response"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
)

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

func (h *Handlers) GetCars(c echo.Context) (err error) {
	var (
		req *request.GetCars
		res *response.GetCars
	)

	// Send response
	defer func() {
		if errResp := c.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	req = request.NewGetCars(c)

	// Call use case
	resUcase, err := h.cars.Get(req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "get cars")
	}

	res = response.NewGetCars(resUcase)
	return nil
}

func (h *Handlers) UpdateCar(c echo.Context) (err error) {
	var (
		req *request.UpdateCar
		res *response.UpdateCar
	)

	// Send response
	defer func() {
		if errResp := c.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewUpdateCar(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call use case
	resUcase, err := h.cars.Update(req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "update car")
	}

	res = response.NewUpdateCar(resUcase)
	return nil
}

func (h *Handlers) DeleteCar(c echo.Context) (err error) {
	var (
		req *request.DeleteCar
		res *response.DeleteCar
	)

	// Send response
	defer func() {
		if errResp := c.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewDeleteCar(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call use case
	resUcase, err := h.cars.Delete(req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "delete car")
	}

	res = response.NewDeleteCar(resUcase)
	return nil
}
