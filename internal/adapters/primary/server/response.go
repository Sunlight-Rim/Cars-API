package server

import (
	"net/http"

	"cars/pkg/errors"

	"github.com/mailru/easyjson"
)

// mapResponse provides a response mapping.
func mapResponse(data any, err error) (int, []byte) {
	var res []byte

	if err != nil {
		status, code, message := errors.GetHTTPErrData(err)

		if res, err = easyjson.Marshal(
			Response{Error: ErrorResponse{Code: code, Message: message}},
		); err != nil {
			return http.StatusInternalServerError, []byte(`{"response": null, "error": {"message": "failed encode response", "code": 0}}`)
		}

		return status, res
	}

	if res, err = easyjson.Marshal(
		Response{Response: data},
	); err != nil {
		return http.StatusInternalServerError, []byte(`{"response": null, "error": {"message": "failed encode response", "code": 0}}`)
	}

	return http.StatusOK, res
}
