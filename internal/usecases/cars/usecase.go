package carsUsecase

import carsDomain "cars/internal/domain/cars"

type usecase struct {
	repo carsDomain.IRepository
}

func New(repo carsDomain.IRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) Create(req *carsDomain.UcaseReqCreate) error
func (u *usecase) Get(req *carsDomain.UcaseReqGet) (*carsDomain.UcaseResGet, error)
func (u *usecase) UpdateColor(req *carsDomain.UcaseReqUpdateColor) error
func (u *usecase) Delete(req *carsDomain.UcaseReqDelete) error
