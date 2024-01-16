package usersRepository

import (
	"database/sql"

	"cars/internal/domain/users"
)

type repository struct {
	psql *sql.DB
}

func New(postgres *sql.DB) *repository {
	return &repository{
		psql: postgres,
	}
}

func (r *repository) Get(req *users.RepoGetReq) (*users.RepoGetRes, error) {
	return nil, nil
}
func (r *repository) UpdateInfo(req *users.RepoUpdateInfoReq) (*users.RepoUpdateInfoRes, error) {
	return nil, nil
}
func (r *repository) UpdatePassword(req *users.RepoUpdatePasswordReq) error { return nil }
func (r *repository) Delete(req *users.RepoDeleteReq) (*users.RepoDeleteRes, error) {
	return nil, nil
}
