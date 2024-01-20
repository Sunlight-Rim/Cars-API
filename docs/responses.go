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

// Returns user ID.
// swagger:response SignupResponse
type SignupResponse struct {
	// in: body
	Body struct {
		Response response.Signup `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns token.
// swagger:response SigninResponse
type SigninResponse struct {
	// in: body
	Body struct {
		Response response.Signin `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns new token.
// swagger:response RefreshResponse
type RefreshResponse struct {
	// in: body
	Body struct {
		Response response.Refresh `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns requested token.
// swagger:response SignoutResponse
type SignoutResponse struct {
	// in: body
	Body struct {
		Response response.Signout `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns revoked tokens.
// swagger:response SignoutAllResponse
type SignoutAllResponse struct {
	// in: body
	Body struct {
		Response response.SignoutAll `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns user account data.
// swagger:response GetMeResponse
type GetMeResponse struct {
	// in: body
	Body struct {
		Response response.GetMe `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns ID of new car.
// swagger:response CreateCarResponse
type CreateCarResponse struct {
	// in: body
	Body struct {
		Response response.CreateCar `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}

// Returns user cars.
// swagger:response GetCarsResponse
type GetCarsResponse struct {
	// in: body
	Body struct {
		Response response.GetCars `json:"response"`
		// example: null
		Error any `json:"error"`
	} `json:"body"`
}
