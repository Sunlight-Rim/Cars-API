package rest

import (
	authEntity "cars/internal/entities/auth"
	carsEntity "cars/internal/entities/cars"
	usersEntity "cars/internal/entities/users"
)

type Service struct {
	auth  authEntity.IUsecase
	users usersEntity.IUsecase
	cars  carsEntity.IUsecase
}

func New(auth authEntity.IUsecase, users usersEntity.IUsecase, cars carsEntity.IUsecase) *Service {
	return &Service{
		auth:  auth,
		users: users,
		cars:  cars,
	}
}
