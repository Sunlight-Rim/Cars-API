package users

import "cars/pkg/errors"

type Usecase struct {
	repo IRepository
}

func New(repo IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

// Get returns user account information.
func (uc *Usecase) Get(req *GetReq) (*GetRes, error) {
	// Get information
	resRepo, err := uc.repo.Get(&RepoGetReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	return NewGetMeRes(resRepo), nil
}

func (uc *Usecase) UpdateInfo(req *UpdateInfoReq) (*UpdateInfoRes, error) {
	return nil, nil
}
func (uc *Usecase) UpdatePassword(req *UpdatePasswordReq) error { return nil }
func (uc *Usecase) Delete(req *DeleteReq) (*DeleteRes, error) {
	return nil, nil
}
