package response

import users "cars/internal/domain/users"

// Get me

type GetMe struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    uint64 `json:"phone"`
}

func NewGetMe(r *users.GetRes) *GetMe {
	return &GetMe{
		ID:       r.ID,
		Username: r.Username,
		Email:    r.Email,
		Phone:    r.Phone,
	}
}
