package response

import "cars/internal/domain/cars"

// Create

type CreateCar struct {
	ID uint64 `json:"id"`
}

func NewCreateCar(r *cars.CreateRes) *CreateCar {
	return &CreateCar{
		ID: r.ID,
	}
}

// Get

type GetCars struct {
	Cars []*cars.Car `json:"cars"`
}

func NewGetCars(r *cars.GetRes) *GetCars {
	return &GetCars{
		Cars: r.Cars,
	}
}
