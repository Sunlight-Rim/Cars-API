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

// swagger:parameters SessionsRequest
type SessionsRequest struct {
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

// swagger:parameters SignoutAllRequest
type SignoutAllRequest struct {
	// in: body
	// required: true
	Body struct {
		Token string `json:"token"`
	}
}

// swagger:parameters UpdateInfoRequest
type UpdateInfoRequest struct {
	// in: body
	// required: true
	Body struct {
		Username string `json:"username"`
		Phone    uint64 `json:"phone"`
	}
}

// swagger:parameters UpdatePasswordRequest
type UpdatePasswordRequest struct {
	// in: body
	// required: true
	Body struct {
		NewPassword string `json:"new_password"`
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

// swagger:parameters UpdateCarRequest
type UpdateCarRequest struct {
	// in: body
	// required: true
	Body struct {
		CarID uint64 `json:"car_id"`
		Model string `json:"model"`
		Color string `json:"color"`
	}
}

// swagger:parameters DeleteCarRequest
type DeleteCarRequest struct {
	// in: body
	// required: true
	Body struct {
		CarID uint64 `json:"car_id"`
	}
}
