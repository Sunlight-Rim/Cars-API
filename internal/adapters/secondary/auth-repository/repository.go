package authRepository

import (
	authEntity "cars/internal/entities/auth"
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

func (r *repository) Signup(*authEntity.RepositoryReqSignup) (*authEntity.RepositoryResSignup, error) {
	return nil, nil
}
func (r *repository) Signin(*authEntity.RepositoryReqSignin) (*authEntity.RepositoryResSignin, error) {
	return nil, nil
}
