package cars

// Create

type CreateReq struct {
	Token string
	Plate string
	Model string
	Color string
}

type CreateRes struct {
	*Car
}

type RepoCreateReq struct {
	Token string
	Plate string
	Model string
	Color string
}

type RepoCreateRes struct {
	*Car
}

// Get

type GetReq struct {
	Token string
}

type GetRes struct {
	Cars []Car
}

type RepoGetReq struct {
	Token string
}

type RepoGetRes struct {
	Cars []Car
}

// Update color

type UpdateColorReq struct {
	Token    string
	CarID    uint64
	NewColor string
}

type UpdateColorRes struct {
	*Car
}

type RepoUpdateColorReq struct {
	Token    string
	CarID    uint64
	NewColor string
}

type RepoUpdateColorRes struct {
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
	Token string
	CarID uint64
}

type RepoDeleteRes struct {
	*Car
}
