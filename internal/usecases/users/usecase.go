package users

import usersEntity "cars/internal/entities/users"

type Usecase struct {
	repo usersEntity.IRepository
}

func New(repo usersEntity.IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (uc *Usecase) Get(req *usersEntity.UsecaseReqGet) (*usersEntity.UsecaseResGet, error) {
	return nil, nil
}
func (uc *Usecase) UpdateInfo(req *usersEntity.UsecaseReqUpdateInfo) (*usersEntity.UsecaseResUpdateInfo, error) {
	return nil, nil
}
func (uc *Usecase) UpdatePassword(req *usersEntity.UsecaseReqUpdatePassword) error { return nil }
func (uc *Usecase) Delete(req *usersEntity.UsecaseReqDelete) (*usersEntity.UsecaseResDelete, error) {
	return nil, nil
}
