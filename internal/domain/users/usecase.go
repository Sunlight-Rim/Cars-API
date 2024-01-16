package users

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

// GetMe parses user ID from token and returns his account information.
func (uc *Usecase) GetMe(req *GetMeReq) (*GetMeRes, error) {
	// Parse token
	claims, err := uc.token.Parse(req.Token)
	if err != nil {
		return nil, errors.Wrap(err, "parse token")
	}

	// Get information
	resRepo, err := uc.repo.GetMe(&RepoGetMeReq{
		ID: claims.UserID,
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
