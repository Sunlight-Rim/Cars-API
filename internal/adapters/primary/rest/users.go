package rest

import (
	"cars/internal/adapters/primary/rest/request"
	"cars/internal/adapters/primary/rest/response"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) GetMe(c echo.Context) (err error) {
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

// UpdateInfo

// UpdatePassword

// DeleteUser
