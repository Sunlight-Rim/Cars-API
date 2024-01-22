package users

import (
	"crypto/sha256"

	"cars/pkg/errors"
)

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

// UpdateInfo changes user information.
func (uc *Usecase) UpdateInfo(req *UpdateInfoReq) (*UpdateInfoRes, error) {
	// Update info
	resRepo, err := uc.repo.UpdateInfo(&RepoUpdateInfoReq{
		UserID:   req.UserID,
		Username: req.Username,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	return NewUpdateInfoRes(resRepo), nil
}

// UpdatePassword changes user password.
func (uc *Usecase) UpdatePassword(req *UpdatePasswordReq) error {
	// Update password
	if err := uc.repo.UpdatePassword(&RepoUpdatePasswordReq{
		UserID:          req.UserID,
		NewPasswordHash: hash(req.NewPassword),
	}); err != nil {
		return errors.Wrap(err, "repository")
	}

	return nil
}

// Delete provides soft delete of user account.
func (uc *Usecase) Delete(req *DeleteReq) (*DeleteRes, error) {
	// Delete info
	resRepo, err := uc.repo.Delete(&RepoDeleteReq{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	return NewDeleteRes(resRepo), nil
}

// hash generates SHA-256 hash string.
func hash(text string) string {
	ph := sha256.New()
	ph.Write([]byte(text)) // No errors
	return string(ph.Sum(nil))
}
