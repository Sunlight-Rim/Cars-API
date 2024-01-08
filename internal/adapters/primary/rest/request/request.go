package request

type Signup struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
