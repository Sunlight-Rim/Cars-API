package auth

import authEntity "cars/internal/entities/auth"

type Usecase struct {
	repo  authEntity.IRepository
	token authEntity.IToken
}

func New(repo authEntity.IRepository, token authEntity.IToken) *Usecase {
	return &Usecase{
		repo:  repo,
		token: token,
	}
}

func (uc *Usecase) Signup(*authEntity.UsecaseReqSignup) (*authEntity.UsecaseResSignup, error) {
	return nil, nil
}
func (uc *Usecase) Signin(*authEntity.UsecaseReqSignin) (*authEntity.UsecaseResSignin, error) {
	return nil, nil
}
func (uc *Usecase) Signout(*authEntity.UsecaseReqSignout) (*authEntity.UsecaseResSignout, error) {
	return nil, nil
}
func (uc *Usecase) Refresh(*authEntity.UsecaseReqRefresh) (*authEntity.UsecaseResRefresh, error) {
	return nil, nil
}
