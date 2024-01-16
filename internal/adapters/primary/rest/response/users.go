package response

import users "cars/internal/domain/users"

// Get me

type GetMe struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

func NewGetMe(r *users.GetMeRes) *GetMe {
	return &GetMe{
		ID:       r.ID,
		Username: r.Username,
		Email:    r.Email,
		Address:  r.Address,
	}
}
