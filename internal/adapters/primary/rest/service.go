package rest

import (
	carsDomain "cars/internal/domain/cars"
	usersDomain "cars/internal/domain/users"
)

type Service struct {
	users usersDomain.IUsecase
	cars  carsDomain.IUsecase
}

func New(users usersDomain.IUsecase, cars carsDomain.IUsecase) *Service {
	return &Service{
		users: users,
		cars:  cars,
	}
}
