package usersUsecase

import usersDomain "cars/internal/domain/users"

type usecase struct {
	repo  usersDomain.IRepository
	token usersDomain.IToken
}

func New(repo usersDomain.IRepository, token usersDomain.IToken) *usecase {
	return &usecase{
		repo:  repo,
		token: token,
	}
}

func (u *usecase) Signup(req *usersDomain.UcaseReqSignup) error { return nil }
func (u *usecase) Signin(req *usersDomain.UcaseReqSignin) (*usersDomain.UcaseResSignin, error) {
	return nil, nil
}
func (u *usecase) Signout(req *usersDomain.UcaseReqSignout) error { return nil }
func (u *usecase) Refresh(req *usersDomain.UcaseReqRefresh) (*usersDomain.UcaseResRefresh, error) {
	return nil, nil
}
func (u *usecase) Get(req *usersDomain.UcaseReqGet) (*usersDomain.UcaseResGet, error) {
	return nil, nil
}
func (u *usecase) UpdateInfo(req *usersDomain.UcaseReqUpdateInfo) error         { return nil }
func (u *usecase) UpdatePassword(req *usersDomain.UcaseReqUpdatePassword) error { return nil }
func (u *usecase) Delete(req *usersDomain.UcaseReqDelete) error                 { return nil }
