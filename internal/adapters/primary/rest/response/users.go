package response

import "cars/internal/domain/users"

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

// Update

type UpdateInfo struct {
	User User `json:"user"`
}

func NewUpdateInfo(r *users.UpdateInfoRes) *UpdateInfo {
	return &UpdateInfo{
		User: User(*r.User),
	}
}

// Delete

type DeleteUser struct {
	User User `json:"user"`
}

func NewDeleteUser(r *users.DeleteRes) *DeleteUser {
	return &DeleteUser{
		User: User(*r.User),
	}
}
