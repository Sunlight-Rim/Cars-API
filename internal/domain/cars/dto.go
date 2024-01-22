package cars

// Create

type CreateReq struct {
	UserID uint64
	Plate  string
	Model  string
	Color  string
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
	UserID uint64
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

func NewGetRes(r *RepoGetRes) *GetRes {
	return &GetRes{Cars: r.Cars}
}

// Update

type UpdateReq struct {
	UserID uint64
	CarID  uint64
	Model  string
	Color  string
}

type UpdateRes struct {
	Car
}

type RepoUpdateReq struct {
	UserID uint64
	CarID  uint64
	Model  string
	Color  string
}

type RepoUpdateRes struct {
	Car
}

func NewUpdateRes(r *RepoUpdateRes) *UpdateRes {
	return &UpdateRes{Car: r.Car}
}

// Delete

type DeleteReq struct {
	UserID uint64
	CarID  uint64
}

type DeleteRes struct {
	Car
}

type RepoDeleteReq struct {
	UserID uint64
	CarID  uint64
}

type RepoDeleteRes struct {
	Car
}

func NewDeleteRes(r *RepoDeleteRes) *DeleteRes {
	return &DeleteRes{Car: r.Car}
}
