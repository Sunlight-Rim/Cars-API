package usersRepository

import (
	usersDomain "cars/internal/domain/users"
	"database/sql"
)

type repository struct {
	psql *sql.DB
}

func New(postgres *sql.DB) *repository {
	return &repository{
		psql: postgres,
	}
}

func (r *repository) Register(req *usersDomain.RepoReqRegister) error
func (r *repository) Login(req *usersDomain.RepoReqLogin) (*usersDomain.RepoResLogin, error)
func (r *repository) Logout(req *usersDomain.RepoReqLogout) error
func (r *repository) Refresh(req *usersDomain.RepoReqRefresh) (*usersDomain.RepoResRefresh, error)
func (r *repository) Get(req *usersDomain.RepoReqGet) (*usersDomain.RepoResGet, error)
func (r *repository) UpdateInfo(req *usersDomain.RepoReqUpdateInfo) error
func (r *repository) UpdatePassword(req *usersDomain.RepoReqUpdatePassword) error
func (r *repository) Delete(req *usersDomain.RepoReqDelete) error
