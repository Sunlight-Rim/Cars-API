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

func (u *usecase) Register(req *usersDomain.UcaseReqRegister) error
func (u *usecase) Login(req *usersDomain.UcaseReqLogin) (*usersDomain.UcaseResLogin, error)
func (u *usecase) Logout(req *usersDomain.UcaseReqLogout) error
func (u *usecase) Refresh(req *usersDomain.UcaseReqRefresh) (*usersDomain.UcaseResRefresh, error)
func (u *usecase) Get(req *usersDomain.UcaseReqGet) (*usersDomain.UcaseResGet, error)
func (u *usecase) UpdateInfo(req *usersDomain.UcaseReqUpdateInfo) error
func (u *usecase) UpdatePassword(req *usersDomain.UcaseReqUpdatePassword) error
func (u *usecase) Delete(req *usersDomain.UcaseReqDelete) error
