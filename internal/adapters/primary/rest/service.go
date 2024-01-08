package rest

import (
	"cars/internal/domain/auth"
	"cars/internal/domain/cars"
	"cars/internal/domain/users"
)

type Service struct {
	auth  auth.IUsecase
	users users.IUsecase
	cars  cars.IUsecase
}

func New(auth auth.IUsecase, users users.IUsecase, cars cars.IUsecase) *Service {
	return &Service{
		auth:  auth,
		users: users,
		cars:  cars,
	}
}
