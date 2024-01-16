package rest

import (
	"cars/internal/adapters/primary/rest/request"
	"cars/internal/adapters/primary/rest/response"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
)

// Signup registers a new user.
func (h *Handlers) Signup(ctx echo.Context) (err error) {
	var (
		req *request.Signup
		res *response.Signup
	)

	// Send response
	defer func() {
		if errResp := ctx.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewSignup(ctx); err != nil {
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

// Signin provides login user to his account.
func (h *Handlers) Signin(ctx echo.Context) (err error) {
	var (
		req *request.Signin
		res *response.Signin
	)

	// Send response
	defer func() {
		if errResp := ctx.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewSignin(ctx); err != nil {
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
