package response

import users "cars/internal/domain/users"

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    uint64 `json:"phone"`
}

// Get me

type GetMe struct {
	User User `json:"user"`
}

func NewGetMe(r *users.GetRes) *GetMe {
	return &GetMe{
		User: User(*r.User),
	}
}
