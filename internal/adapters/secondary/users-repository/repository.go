package usersRepository

import (
	usersEntity "cars/internal/entities/users"
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

func (r *repository) Get(req *usersEntity.RepositoryReqGet) (*usersEntity.RepositoryResGet, error) {
	return nil, nil
}
func (r *repository) UpdateInfo(req *usersEntity.RepositoryReqUpdateInfo) (*usersEntity.RepositoryResUpdateInfo, error) {
	return nil, nil
}
func (r *repository) UpdatePassword(req *usersEntity.RepositoryReqUpdatePassword) error { return nil }
func (r *repository) Delete(req *usersEntity.RepositoryReqDelete) (*usersEntity.RepositoryResDelete, error) {
	return nil, nil
}
