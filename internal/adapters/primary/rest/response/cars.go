package response

import "cars/internal/domain/cars"

type CreateCar struct {
	ID uint64 `json:"id"`
}

func NewCreateCar(r *cars.CreateRes) *CreateCar {
	return &CreateCar{
		ID: r.ID,
	}
}
