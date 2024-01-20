package response

import (
	"cars/internal/domain/auth"
)

// General response struct.
type Response struct {
	Response any `json:"response"`
	Error    any `json:"error"`
}

// General error response struct.
type ErrorResponse struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

// Signup

type Signup struct {
	ID uint64 `json:"id"`
}

func NewSignup(r *auth.SignupRes) *Signup {
	return &Signup{ID: r.ID}
}

// Signin

type Signin struct {
	Token string `json:"token"`
}

func NewSignin(r *auth.SigninRes) *Signin {
	return &Signin{Token: r.Token}
}

// Refresh

type Refresh struct {
	Token string `json:"token"`
}

func NewRefresh(r *auth.RefreshRes) *Refresh {
	return &Refresh{Token: r.Token}
}

// Signout

// Singout all
