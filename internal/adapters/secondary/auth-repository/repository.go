package authRepository

import (
	"cars/internal/domain/auth"
	"cars/pkg/errors"
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

func (r *repository) Signup(req *auth.RepoSignupReq) (*auth.RepoSignupRes, error) {
	// Start transacton
	tx, err := r.psql.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "begin transaction")
	}

	// End transacton
	defer func() {
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				err = errors.Wrapf(errRollback, "rollback, %v", err)
			}
		} else {
			if errCommit := tx.Commit(); errCommit != nil {
				err = errors.Wrap(errCommit, "commit")
			}
		}
	}()

	// Check if email is busy
	var emailBusy bool

	if err := tx.QueryRow(`
		SELECT true
		FROM api.users
		WHERE email = $1
	`, req.Email,
	).Scan(&emailBusy); err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "checking email")
	}

	if emailBusy {
		return nil, errors.Wrap(errors.EmailIsBusy, "busy email")
	}

	// Add user
	var id uint64

	if err := tx.QueryRow(`
		INSERT INTO api.users(
			username,
			email,
			"address",
			password_hash
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		req.Username,
		req.Email,
		req.Address,
		req.PasswordHash,
	).Scan(&id); err != nil {
		return nil, errors.Wrap(err, "adding user")
	}

	return &auth.RepoSignupRes{ID: id}, nil
}

func (r *repository) Signin(*auth.RepoSigninReq) (*auth.RepoSigninRes, error) {
	return nil, nil
}
