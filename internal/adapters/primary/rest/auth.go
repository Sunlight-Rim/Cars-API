package rest

import (
	"cars/internal/adapters/primary/rest/request"
	"cars/internal/adapters/primary/rest/response"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) Signup(c echo.Context) (err error) {
	var (
		req *request.Signup
		res *response.Signup
	)

	// Send response
	defer func() {
		if errResp := c.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewSignup(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	resUcase, err := h.auth.Signup(req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "signup")
	}

	res = response.NewSignup(resUcase)
	return nil
}

func (h *Handlers) Signin(c echo.Context) (err error) {
	var (
		req *request.Signin
		res *response.Signin
	)

	// Send response
	defer func() {
		if errResp := c.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewSignin(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	resUcase, err := h.auth.Signin(req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "signin")
	}

	res = response.NewSignin(resUcase)
	return nil
}

func (h *Handlers) Refresh(c echo.Context) (err error) {
	var (
		req *request.Refresh
		res *response.Refresh
	)

	// Send response
	defer func() {
		if errResp := c.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewRefresh(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	resUcase, err := h.auth.Refresh(req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "refresh")
	}

	res = response.NewRefresh(resUcase)
	return nil
}

func (h *Handlers) Signout(c echo.Context) (err error) {
	var (
		req *request.Signout
		res *response.Signout
	)

	// Send response
	defer func() {
		if errResp := c.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewSignout(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	resUcase, err := h.auth.Signout(req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "signout")
	}

	res = response.NewSignout(resUcase)
	return nil
}
