package auth

import (
	"crypto/sha256"

	"cars/pkg/errors"
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
	resRepo, err := uc.repo.Signup(&RepoSignupReq{
		Username:     req.Username,
		Email:        req.Email,
		Address:      req.Address,
		PasswordHash: string(ph.Sum(nil)),
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	return NewSignupRes(resRepo), nil
}

// Signin checks credentials, generates token and saves them to cache.
func (uc *Usecase) Signin(req *SigninReq) (*SigninRes, error) {
	// Hash password
	ph := sha256.New()
	if _, err := ph.Write([]byte(req.Password)); err != nil {
		return nil, errors.Wrap(err, "hash password")
	}

	// Check credentials
	resRepo, err := uc.repo.Signin(&RepoSigninReq{
		Email:        req.Email,
		PasswordHash: string(ph.Sum(nil)),
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository")
	}

	// Create token
	token, err := uc.token.Create(resRepo.ID)
	if err != nil {
		return nil, errors.Wrap(err, "create token")
	}

	// Save token
	if err := uc.token.SaveRefresh(token, resRepo.ID); err != nil {
		return nil, errors.Wrap(err, "save token")
	}

	return NewSigninRes(token), nil
}

func (uc *Usecase) Refresh(*RefreshReq) (*RefreshRes, error) {
	return nil, nil
}

func (uc *Usecase) Signout(*SignoutReq) (*SignoutRes, error) {
	return nil, nil
}

func (uc *Usecase) SignoutAll(*SignoutAllReq) (*SignoutAllRes, error) {
	return nil, nil
}
