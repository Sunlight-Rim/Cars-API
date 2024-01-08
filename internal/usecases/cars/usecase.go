package cars

import carsEntity "cars/internal/entities/cars"

type Usecase struct {
	repo carsEntity.IRepository
}

func New(repo carsEntity.IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (uc *Usecase) Create(req *carsEntity.UsecaseReqCreate) (*carsEntity.UsecaseResCreate, error) {
	return nil, nil
}
func (uc *Usecase) Get(req *carsEntity.UsecaseReqGet) (*carsEntity.UsecaseResGet, error) {
	return nil, nil
}
func (uc *Usecase) UpdateColor(req *carsEntity.UsecaseReqUpdateColor) (*carsEntity.UsecaseResUpdateColor, error) {
	return nil, nil
}
func (uc *Usecase) Delete(req *carsEntity.UsecaseReqDelete) (*carsEntity.UsecaseResDelete, error) {
	return nil, nil
}
