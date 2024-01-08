package rest

import (
	"cars/internal/adapters/primary/rest/request"
	"cars/internal/adapters/primary/rest/response"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
)

func (s *Service) Signup(ctx echo.Context) (err error) {
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
	id, err := s.auth.Signup(req.ToEntity())
	if err != nil {
		return errors.Wrap(err, "signup")
	}

	res = response.NewSignup(id)
	return nil
}
