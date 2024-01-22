package users

// Get

type GetReq struct {
	UserID uint64
}

type GetRes struct {
	*User
}

type RepoGetReq struct {
	UserID uint64
}

type RepoGetRes struct {
	*User
}

func NewGetMeRes(r *RepoGetRes) *GetRes {
	return &GetRes{User: r.User}
}

// Update info

type UpdateInfoReq struct {
	UserID   uint64
	Username string
	Phone    uint64
}

type UpdateInfoRes struct {
	*User
}

type RepoUpdateInfoReq struct {
	UserID   uint64
	Username string
	Phone    uint64
}

type RepoUpdateInfoRes struct {
	*User
}

func NewUpdateInfoRes(r *RepoUpdateInfoRes) *UpdateInfoRes {
	return &UpdateInfoRes{User: r.User}
}

// Update password

type UpdatePasswordReq struct {
	UserID      uint64
	NewPassword string
}

type RepoUpdatePasswordReq struct {
	UserID          uint64
	NewPasswordHash string
}

// Delete

type DeleteReq struct {
	UserID uint64
}

type DeleteRes struct {
	*User
}

type RepoDeleteReq struct {
	UserID uint64
}

type RepoDeleteRes struct {
	*User
}

func NewDeleteRes(r *RepoDeleteRes) *DeleteRes {
	return &DeleteRes{User: r.User}
}
