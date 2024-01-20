package cars

// Create

type CreateReq struct {
	Token string
	Plate string
	Model string
	Color string
}

type CreateRes struct {
	ID uint64
}

type RepoCreateReq struct {
	UserID uint64
	Plate  string
	Model  string
	Color  string
}

type RepoCreateRes struct {
	ID uint64
}

func NewCreate(r *RepoCreateRes) *CreateRes {
	return &CreateRes{ID: r.ID}
}

// Get

type GetReq struct {
	Token string
}

type GetRes struct {
	Cars []*Car
}

type RepoGetReq struct {
	UserID uint64
}

type RepoGetRes struct {
	Cars []*Car
}

// Update

type UpdateReq struct {
	Token string
	CarID uint64
	Model string
	Color string
}

type UpdateRes struct {
	*Car
}

type RepoUpdateReq struct {
	UserID uint64
	CarID  uint64
	Model  string
	Color  string
}

type RepoUpdateRes struct {
	*Car
}

// Delete

type DeleteReq struct {
	Token string
	CarID uint64
}

type DeleteRes struct {
	*Car
}

type RepoDeleteReq struct {
	UserID uint64
	CarID  uint64
}

type RepoDeleteRes struct {
	*Car
}
