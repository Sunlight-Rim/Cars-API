package users

// Get me

type GetMeReq struct {
	Token string
}

type GetMeRes struct {
	*User
}

type RepoGetMeReq struct {
	ID uint64
}

type RepoGetMeRes struct {
	*User
}

func NewGetMeRes(r *RepoGetMeRes) *GetMeRes {
	return &GetMeRes{User: r.User}
}

// Update info

type UpdateInfoReq struct {
	Token    string
	Username string
	Address  string
}

type UpdateInfoRes struct {
	*User
}

type RepoUpdateInfoReq struct {
	ID       uint64
	Username string
	Address  string
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
