package usersRepository

import (
	"database/sql"

	"cars/internal/domain/users"
	"cars/pkg/errors"
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
	var user users.User

	// Add user
	if err := r.psql.QueryRow(`
		SELECT
			id,
			username,
			email,
			phone
		FROM api.users
		WHERE id = $1
	`, req.ID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Phone,
	); err != nil {
		return nil, errors.Wrap(err, "adding user")
	}

	return &users.RepoGetRes{User: &user}, nil
}

func (r *repository) UpdateInfo(req *users.RepoUpdateInfoReq) (*users.RepoUpdateInfoRes, error) {
	return nil, nil
}
func (r *repository) UpdatePassword(req *users.RepoUpdatePasswordReq) error { return nil }
func (r *repository) Delete(req *users.RepoDeleteReq) (*users.RepoDeleteRes, error) {
	return nil, nil
}
