package cars

import "cars/pkg/errors"

type Usecase struct {
	repo IRepository
}

func New(repo IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

// Create adds car to user in repository.
func (uc *Usecase) Create(req *CreateReq) (*CreateRes, error) {
	// Add car to user in database
	resRepo, err := uc.repo.Create(&RepoCreateReq{
		UserID: req.UserID,
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
	// Add car to user in database
	resRepo, err := uc.repo.Get(&RepoGetReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	return NewGetRes(resRepo), nil
}

// Update changes user car in repository and returns changed car.
func (uc *Usecase) Update(req *UpdateReq) (*UpdateRes, error) {
	// Add car to user in database
	resRepo, err := uc.repo.Update(&RepoUpdateReq{
		UserID: req.UserID,
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
	// Add car to user in database
	resRepo, err := uc.repo.Delete(&RepoDeleteReq{
		UserID: req.UserID,
		CarID:  req.CarID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	return NewDeleteRes(resRepo), nil
}
