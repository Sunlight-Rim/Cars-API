package response

import "cars/internal/domain/cars"

type Car struct {
	ID    uint64 `json:"id"`
	Plate string `json:"plate"`
	Model string `json:"model"`
	Color string `json:"color"`
}

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
	Cars []Car `json:"cars"`
}

func NewGetCars(r *cars.GetRes) *GetCars {
	var res GetCars

	for i := range r.Cars {
		res.Cars = append(res.Cars, Car(*r.Cars[i]))
	}

	return &res
}

// Update

type UpdateCar struct {
	Car Car `json:"car"`
}

func NewUpdateCar(r *cars.UpdateRes) *UpdateCar {
	return &UpdateCar{
		Car: Car(r.Car),
	}
}

// Delete

type DeleteCar struct {
	Car Car `json:"car"`
}

func NewDeleteCar(r *cars.DeleteRes) *DeleteCar {
	return &DeleteCar{
		Car: Car(r.Car),
	}
}
