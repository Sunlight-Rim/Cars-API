// Code generated, DO NOT EDIT.
package errors

import "net/http"

const unknown = "unknown"

// Example usage for output errors.response.json
// echo.GET("/errors", func(ctx echo.Context) error {
//   return ctx.JSONBlob(http.StatusOK, errors.ResponseList)
// })
var ResponseList []byte = []byte(`{
    "1447728996": {
        "en": "Invalid request content"
    },
    "1760844835": {
        "en": "Expired token"
    },
    "2213895429": {
        "en": "Method not allowed"
    },
    "2352452421": {
        "en": "Invalid token"
    },
    "3232214746": {
        "en": "Not found"
    },
    "3233457132": {
        "en": "Invalid request format"
    },
    "4129226438": {
        "en": "Plate is busy"
    },
    "683930722": {
        "en": "Email is busy"
    },
    "816852758": {
        "en": "Invalid credentials"
    }
}`) // `

// Error codes
const (
	CodeNotFound              uint32 = 3232214746
	CodeMethodNotAllowed      uint32 = 2213895429
	CodeInvalidRequestFormat  uint32 = 3233457132
	CodeInvalidRequestContent uint32 = 1447728996
	CodeEmailIsBusy           uint32 = 683930722
	CodePlateIsBusy           uint32 = 4129226438
	CodeInvalidCredentials    uint32 = 816852758
	CodeInvalidToken          uint32 = 2352452421
	CodeExpiredToken          uint32 = 1760844835
)

// Error Variables
var (
	NotFound              error = New("not found", CodeNotFound)
	MethodNotAllowed      error = New("method not allowed", CodeMethodNotAllowed)
	InvalidRequestFormat  error = New("invalid request format", CodeInvalidRequestFormat)
	InvalidRequestContent error = New("invalid request content", CodeInvalidRequestContent)
	EmailIsBusy           error = New("email is busy", CodeEmailIsBusy)
	PlateIsBusy           error = New("plate is busy", CodePlateIsBusy)
	InvalidCredentials    error = New("invalid credentials", CodeInvalidCredentials)
	InvalidToken          error = New("invalid token", CodeInvalidToken)
	ExpiredToken          error = New("expired token", CodeExpiredToken)
)

// Hash map data by error codes
var HttpResponse = map[uint32]map[string]any{
	3232214746: {"status": 404, "text": "Not found"},
	2213895429: {"status": 405, "text": "Method not allowed"},
	3233457132: {"status": 400, "text": "Invalid request format"},
	1447728996: {"status": 422, "text": "Invalid request content"},
	683930722:  {"status": 422, "text": "Email is busy"},
	4129226438: {"status": 422, "text": "Plate is busy"},
	816852758:  {"status": 401, "text": "Invalid credentials"},
	2352452421: {"status": 401, "text": "Invalid token"},
	1760844835: {"status": 401, "text": "Expired token"},
}

// GetHTTPErrData returning http status, error message and error code
func GetHTTPErrData(err error) (int, uint32, string) {
	code, _ := GetCode(err)

	return getHttpStatus(code), code, getErrText(code)
}

func getErrText(code uint32) string {
	if _, ok := HttpResponse[code]; !ok {
		return unknown
	}

	if _, ok := HttpResponse[code]["text"]; !ok {
		return unknown
	}

	if status, ok := HttpResponse[code]["text"].(string); ok {
		return status
	}

	return unknown
}

func getHttpStatus(code uint32) int {
	if _, ok := HttpResponse[code]; !ok {
		return http.StatusInternalServerError
	}

	if _, ok := HttpResponse[code]["status"]; !ok {
		return http.StatusInternalServerError
	}

	if status, ok := HttpResponse[code]["status"].(int); ok {
		return status
	}

	return http.StatusInternalServerError
}
