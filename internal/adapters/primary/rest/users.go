package rest

import (
	"cars/internal/adapters/primary/rest/request"
	"cars/internal/adapters/primary/rest/response"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) getMe(c echo.Context) (err error) {
	var (
		req *request.GetMe
		res *response.GetMe
	)

	// Send response
	defer func() {
		if errResp := c.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewGetMe(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	resUcase, err := h.users.Get(req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "get")
	}

	res = response.NewGetMe(resUcase)
	return nil
}

func (h *Handlers) updateInfo(c echo.Context) (err error) {
	var (
		req *request.UpdateInfo
		res *response.UpdateInfo
	)

	// Send response
	defer func() {
		if errResp := c.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewUpdateInfo(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	resUcase, err := h.users.UpdateInfo(req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "update info")
	}

	res = response.NewUpdateInfo(resUcase)
	return nil
}

func (h *Handlers) updatePassword(c echo.Context) (err error) {
	var req *request.UpdatePassword

	// Send response
	defer func() {
		if errResp := c.JSONBlob(response.Map(nil, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewUpdatePassword(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	if err := h.users.UpdatePassword(req.ToEntity()); err != nil {
		return errors.Wrap(err, "update password")
	}

	return nil
}

func (h *Handlers) deleteUser(c echo.Context) (err error) {
	var (
		req *request.DeleteUser
		res *response.DeleteUser
	)

	// Send response
	defer func() {
		if errResp := c.JSONBlob(response.Map(res, err)); errResp != nil {
			err = errors.Wrapf(errResp, "response, %v", err)
		}
	}()

	// Parse request
	if req, err = request.NewDeleteUser(c); err != nil {
		return errors.Wrap(err, "request")
	}

	// Call usecase
	resUcase, err := h.users.Delete(req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "delete user")
	}

	res = response.NewDeleteUser(resUcase)
	return nil
}
