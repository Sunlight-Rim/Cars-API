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
	var res = users.RepoGetRes{User: &users.User{
		ID: req.UserID,
	}}

	if err := r.psql.QueryRow(`
		SELECT
			username,
			email,
			phone
		FROM api.users
		WHERE id = $1
	`, req.UserID).Scan(
		&res.Username,
		&res.Email,
		&res.Phone,
	); err != nil {
		return nil, errors.Wrap(err, "getting user")
	}

	return &res, nil
}

func (r *repository) UpdateInfo(req *users.RepoUpdateInfoReq) (*users.RepoUpdateInfoRes, error) {
	var res = users.RepoUpdateInfoRes{User: &users.User{
		ID:       req.UserID,
		Username: req.Username,
		Phone:    req.Phone,
	}}

	if err := r.psql.QueryRow(`
		UPDATE api.users
		SET
			username = $1,
			phone = $2
		WHERE id = $3
		RETURNING email
	`,
		req.Username,
		req.Phone,
		req.UserID,
	).Scan(res.Email); err != nil {
		return nil, errors.Wrap(err, "updating user")
	}

	return &res, nil
}

func (r *repository) UpdatePassword(req *users.RepoUpdatePasswordReq) error {
	if _, err := r.psql.Exec(`
		UPDATE api.users
		SET password_hash = $1
		WHERE id = $2
	`,
		req.NewPasswordHash,
		req.UserID,
	); err != nil {
		return errors.Wrap(err, "updating password")
	}

	return nil
}

func (r *repository) Delete(req *users.RepoDeleteReq) (*users.RepoDeleteRes, error) {
	var res = users.RepoDeleteRes{User: &users.User{
		ID: req.UserID,
	}}

	if err := r.psql.QueryRow(`
		DELETE api.users
		WHERE id = $1
		RETURNING
			username,
			email,
			phone
	`, req.UserID).Scan(
		res.Username,
		res.Email,
		res.Phone,
	); err != nil {
		return nil, errors.Wrap(err, "deleting user")
	}

	return &res, nil
}
