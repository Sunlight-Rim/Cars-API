package response

import (
	"net/http"

	"cars/pkg/errors"

	"github.com/mailru/easyjson"
)

var failedResponse = []byte(`{"response": null, "error": {"message": "failed encode response", "code": 0}}`)

// Map provides a response mapping.
func Map(data any, err error) (int, []byte) {
	var res []byte

	if err != nil {
		status, code, message := errors.GetHTTPErrData(err)

		if res, err = easyjson.Marshal(
			Response{Error: ErrorResponse{Code: code, Message: message}},
		); err != nil {
			return http.StatusInternalServerError, failedResponse
		}

		return status, res
	}

	if res, err = easyjson.Marshal(
		Response{Response: data},
	); err != nil {
		return http.StatusInternalServerError, failedResponse
	}

	return http.StatusOK, res
}
