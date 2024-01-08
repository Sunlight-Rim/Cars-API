package response

import "cars/internal/domain/auth"

// General response struct.
type Response struct {
	Response any           `json:"response"`
	Error    ErrorResponse `json:"error"`
}

// General error response struct.
type ErrorResponse struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

// Signup

type Signup struct {
	ID uint64
}

func NewSignup(r *auth.SignupRes) *Signup {
	return &Signup{ID: r.ID}
}
