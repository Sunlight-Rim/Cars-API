package carsDomain

// Usecase

type UcaseReqCreate struct {
	Token string
	Plate string
	Model string
	Color string
}

type UcaseReqGet struct {
	Token string
}

type UcaseResGet struct {
	Cars []Car
}

type UcaseReqUpdateColor struct {
	Token    string
	CarID    uint64
	NewColor string
}

type UcaseReqDelete struct {
	Token string
	CarID uint64
}

// Repository

type RepoReqCreate struct {
	Token string
	Plate string
	Model string
	Color string
}

type RepoReqGet struct {
	Token string
}

type RepoResGet struct {
	Cars []Car
}

type RepoReqUpdateColor struct {
	Token    string
	CarID    uint64
	NewColor string
}

type RepoReqDelete struct {
	Token string
	CarID uint64
}
