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