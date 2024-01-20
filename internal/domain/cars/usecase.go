package cars

type Usecase struct {
	repo  IRepository
	token IToken
}

func New(repo IRepository, token IToken) *Usecase {
	return &Usecase{
		repo:  repo,
		token: token,
	}
}

func (uc *Usecase) Create(req *CreateReq) (*CreateRes, error) {
	return nil, nil
}
func (uc *Usecase) Get(req *GetReq) (*GetRes, error) {
	return nil, nil
}
func (uc *Usecase) Update(req *UpdateReq) (*UpdateRes, error) {
	return nil, nil
}
func (uc *Usecase) Delete(req *DeleteReq) (*DeleteRes, error) {
	return nil, nil
}
