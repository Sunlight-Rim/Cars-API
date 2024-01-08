package users

// Get

type GetReq struct {
	ID uint64
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

// Update info

type UpdateInfoReq struct {
	ID       uint64
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
	ID          uint64
	NewPassword string
}

type RepoUpdatePasswordReq struct {
	ID          uint64
	NewPassword string
}

// Delete

type DeleteReq struct {
	ID uint64
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
