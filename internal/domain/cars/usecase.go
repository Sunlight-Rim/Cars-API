package cars

import "cars/pkg/errors"

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

// Create adds car to user in database.
func (uc *Usecase) Create(req *CreateReq) (*CreateRes, error) {
	// Parse token
	claims, err := uc.token.Parse(req.Token)
	if err != nil {
		return nil, errors.Wrap(err, "parse token")
	}

	// Add car to user in database
	resRepo, err := uc.repo.Create(&RepoCreateReq{
		UserID: claims.UserID,
		Plate:  req.Plate,
		Model:  req.Model,
		Color:  req.Color,
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	return NewCreate(resRepo), nil
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
