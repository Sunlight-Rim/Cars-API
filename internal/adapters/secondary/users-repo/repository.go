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

func (r *repository) Signup(req *usersDomain.RepoReqSignup) error { return nil }
func (r *repository) Signin(req *usersDomain.RepoReqSignin) (*usersDomain.RepoResSignin, error) {
	return nil, nil
}
func (r *repository) Signout(req *usersDomain.RepoReqSignout) error { return nil }
func (r *repository) Get(req *usersDomain.RepoReqGet) (*usersDomain.RepoResGet, error) {
	return nil, nil
}
func (r *repository) UpdateInfo(req *usersDomain.RepoReqUpdateInfo) error         { return nil }
func (r *repository) UpdatePassword(req *usersDomain.RepoReqUpdatePassword) error { return nil }
func (r *repository) Delete(req *usersDomain.RepoReqDelete) error                 { return nil }
