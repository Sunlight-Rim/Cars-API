package carsEntity

// Create

type RepositoryReqCreate struct {
	Token string
	Plate string
	Model string
	Color string
}

type RepositoryResCreate struct {
	*Car
}

// Get

type RepositoryReqGet struct {
	Token string
}

type RepositoryResGet struct {
	Cars []Car
}

// Update color

type RepositoryReqUpdateColor struct {
	Token    string
	CarID    uint64
	NewColor string
}

type RepositoryResUpdateColor struct {
	*Car
}

// Delete

type RepositoryReqDelete struct {
	Token string
	CarID uint64
}

type RepositoryResDelete struct {
	*Car
}
