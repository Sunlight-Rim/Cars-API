package usersEntity

// Get

type UsecaseReqGet struct {
	ID uint64
}

type UsecaseResGet struct {
	*User
}

// Update info

type UsecaseReqUpdateInfo struct {
	ID       uint64
	Username string
	Address  string
}

type UsecaseResUpdateInfo struct {
	*User
}

// Update password

type UsecaseReqUpdatePassword struct {
	ID          uint64
	NewPassword string
}

// Delete

type UsecaseReqDelete struct {
	ID uint64
}

type UsecaseResDelete struct {
	*User
}
