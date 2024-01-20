package model

// swagger:parameters SignupRequest
type SignupRequest struct {
	// in: body
	// required: true
	Body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    uint64 `json:"phone"`
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

// swagger:parameters RefreshRequest
type RefreshRequest struct {
	// in: body
	// required: true
	Body struct {
		Token string `json:"token"`
	}
}

// swagger:parameters SignoutRequest
type SignoutRequest struct {
	// in: body
	// required: true
	Body struct {
		Token string `json:"token"`
	}
}

// swagger:parameters CreateCarRequest
type CreateCarRequest struct {
	// in: body
	// required: true
	Body struct {
		Plate string `json:"plate"`
		Model string `json:"model"`
		Color string `json:"color"`
	}
}
