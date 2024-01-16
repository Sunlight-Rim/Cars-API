package model

// swagger:parameters SignupRequest
type SignupRequest struct {
	// in: body
	// required: true
	Body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Address  string `json:"address"`
		Password string `json:"password"`
	}
}

// swagger:parameters SigninRequest
type SigninRequest struct {
	// in: body
	// required: true
	Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}
