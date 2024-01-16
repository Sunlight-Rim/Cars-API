package model

import "cars/internal/adapters/primary/rest/response"

// General errors model. Response field is null.
// swagger:response ErrorResponse
type ErrorResponse struct {
	// in: body
	Body struct {
		// example: null
		Response any                    `json:"response"`
		Error    response.ErrorResponse `json:"error"`
	} `json:"body"`
}

// A list of errors containing error codes and text descriptions.
// swagger:response ErrorsListResponse
type ErrorsListResponse struct {
	// in: body
	Body struct {
		Response struct {
			Language string `json:"language"`
		} `json:"error_code"`
	} `json:"body"`
}

// Successfully registered.
// swagger:response SignupResponse
type SignupResponse struct {
	// in: body
	Body struct {
		Response response.Signup `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Successfully signed in. Returns token.
// swagger:response SigninResponse
type SigninResponse struct {
	// in: body
	Body struct {
		Response response.Signin `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}
