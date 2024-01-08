package cars

type Usecase struct {
	repo IRepository
}

func New(repo IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (uc *Usecase) Create(req *CreateReq) (*CreateRes, error) {
	return nil, nil
}
func (uc *Usecase) Get(req *GetReq) (*GetRes, error) {
	return nil, nil
}
func (uc *Usecase) UpdateColor(req *UpdateColorReq) (*UpdateColorRes, error) {
	return nil, nil
}
func (uc *Usecase) Delete(req *DeleteReq) (*DeleteRes, error) {
	return nil, nil
}
