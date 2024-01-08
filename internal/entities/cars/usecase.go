package carsEntity

// Create

type UsecaseReqCreate struct {
	Token string
	Plate string
	Model string
	Color string
}

type UsecaseResCreate struct {
	*Car
}

// Get

type UsecaseReqGet struct {
	Token string
}

type UsecaseResGet struct {
	Cars []Car
}

// Update color

type UsecaseReqUpdateColor struct {
	Token    string
	CarID    uint64
	NewColor string
}

type UsecaseResUpdateColor struct {
	*Car
}

// Delete

type UsecaseReqDelete struct {
	Token string
	CarID uint64
}

type UsecaseResDelete struct {
	*Car
}
