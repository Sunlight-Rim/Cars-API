package auth

import (
	"cars/pkg/errors"
	"crypto/sha256"
)

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

// Signup creates a new user.
func (uc *Usecase) Signup(req *SignupReq) (*SignupRes, error) {
	// Hash password
	ph := sha256.New()
	if _, err := ph.Write([]byte(req.Password)); err != nil {
		return nil, errors.Wrap(err, "hash password")
	}

	// Add user to database
	res, err := uc.repo.Signup(&RepoSignupReq{
		Username:     req.Username,
		Email:        req.Email,
		Address:      req.Address,
		PasswordHash: string(ph.Sum(nil)),
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	return &SignupRes{ID: res.ID}, nil
}

func (uc *Usecase) Signin(*SigninReq) (*SigninRes, error) {
	return nil, nil
}
func (uc *Usecase) Signout(*SignoutReq) (*SignoutRes, error) {
	return nil, nil
}
func (uc *Usecase) Refresh(*RefreshReq) (*RefreshRes, error) {
	return nil, nil
}
