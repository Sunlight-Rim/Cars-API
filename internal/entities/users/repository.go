package usersEntity

// Get

type RepositoryReqGet struct {
	ID uint64
}

type RepositoryResGet struct {
	*User
}

// Update info

type RepositoryReqUpdateInfo struct {
	ID       uint64
	Username string
	Address  string
}

type RepositoryResUpdateInfo struct {
	*User
}

// Update password

type RepositoryReqUpdatePassword struct {
	ID          uint64
	NewPassword string
}

// Delete

type RepositoryReqDelete struct {
	ID uint64
}

type RepositoryResDelete struct {
	*User
}
