package response

import "cars/internal/domain/auth"

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
	UserID uint64 `json:"id"`
}

func NewSignup(r *auth.SignupRes) *Signup {
	return &Signup{UserID: r.UserID}
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

// Sessions

type Sessions struct {
	Tokens []string `json:"tokens"`
}

func NewSessions(r *auth.SessionsRes) *Sessions {
	return &Sessions{Tokens: r.Tokens}
}

// Signout

type Signout struct {
	Token string `json:"token"`
}

func NewSignout(r *auth.SignoutRes) *Signout {
	return &Signout{Token: r.Token}
}

// Singout all

type SignoutAll struct {
	Tokens []string `json:"tokens"`
}

func NewSignoutAll(r *auth.SignoutAllRes) *SignoutAll {
	return &SignoutAll{Tokens: r.Tokens}
}
