package users

// Get

type GetReq struct {
	Token string
}

type GetRes struct {
	*User
}

type RepoGetReq struct {
	ID uint64
}

type RepoGetRes struct {
	*User
}

func NewGetMeRes(r *RepoGetRes) *GetRes {
	return &GetRes{User: r.User}
}

// Update info

type UpdateInfoReq struct {
	Token    string
	Username string
	Phone    uint64
}

type UpdateInfoRes struct {
	*User
}

type RepoUpdateInfoReq struct {
	ID       uint64
	Username string
	Phone    uint64
}

type RepoUpdateInfoRes struct {
	*User
}

// Update password

type UpdatePasswordReq struct {
	Token       string
	NewPassword string
}

type RepoUpdatePasswordReq struct {
	ID              uint64
	NewPasswordHash string
}

// Delete

type DeleteReq struct {
	Token string
}

type DeleteRes struct {
	*User
}

type RepoDeleteReq struct {
	ID uint64
}

type RepoDeleteRes struct {
	*User
}
