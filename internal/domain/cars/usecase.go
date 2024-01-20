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

// Create adds car to user in repository.
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

// Get returns all user cars from repository.
func (uc *Usecase) Get(req *GetReq) (*GetRes, error) {
	// Parse token
	claims, err := uc.token.Parse(req.Token)
	if err != nil {
		return nil, errors.Wrap(err, "parse token")
	}

	// Add car to user in database
	resRepo, err := uc.repo.Get(&RepoGetReq{
		UserID: claims.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	return NewGetRes(resRepo), nil
}

// Update changes user car in repository and returns changed car.
func (uc *Usecase) Update(req *UpdateReq) (*UpdateRes, error) {
	// Parse token
	claims, err := uc.token.Parse(req.Token)
	if err != nil {
		return nil, errors.Wrap(err, "parse token")
	}

	// Add car to user in database
	resRepo, err := uc.repo.Update(&RepoUpdateReq{
		UserID: claims.UserID,
		CarID:  req.CarID,
		Model:  req.Model,
		Color:  req.Color,
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	return NewUpdateRes(resRepo), nil
}

// Delete deletes user car in repository.
func (uc *Usecase) Delete(req *DeleteReq) (*DeleteRes, error) {
	// Parse token
	claims, err := uc.token.Parse(req.Token)
	if err != nil {
		return nil, errors.Wrap(err, "parse token")
	}

	// Add car to user in database
	resRepo, err := uc.repo.Delete(&RepoDeleteReq{
		UserID: claims.UserID,
		CarID:  req.CarID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	return NewDeleteRes(resRepo), nil
}
