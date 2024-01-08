package users

type Usecase struct {
	repo IRepository
}

func New(repo IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (uc *Usecase) Get(req *GetReq) (*GetRes, error) {
	return nil, nil
}
func (uc *Usecase) UpdateInfo(req *UpdateInfoReq) (*UpdateInfoRes, error) {
	return nil, nil
}
func (uc *Usecase) UpdatePassword(req *UpdatePasswordReq) error { return nil }
func (uc *Usecase) Delete(req *DeleteReq) (*DeleteRes, error) {
	return nil, nil
}
