package authRepository

import (
	"database/sql"

	"cars/internal/domain/auth"
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

func (r *repository) Signup(req *auth.RepoSignupReq) (_ *auth.RepoSignupRes, err error) {
	var res auth.RepoSignupRes

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
	).Scan(&emailBusy); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, "checking email")
	}

	if emailBusy {
		return nil, errors.Wrap(errors.EmailIsBusy, "busy email")
	}

	// Add user
	if err := tx.QueryRow(`
		INSERT INTO api.users(
			username,
			email,
			phone,
			password_hash
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		req.Username,
		req.Email,
		req.Phone,
		req.PasswordHash,
	).Scan(&res.ID); err != nil {
		return nil, errors.Wrap(err, "adding user")
	}

	return &res, nil
}

func (r *repository) Signin(req *auth.RepoSigninReq) (*auth.RepoSigninRes, error) {
	var res auth.RepoSigninRes

	if err := r.psql.QueryRow(`
		SELECT id
		FROM api.users
		WHERE
			email = $1 AND
			password_hash = $2 AND
			removed = FALSE
	`,
		req.Email,
		req.PasswordHash,
	).Scan(&res.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(errors.InvalidCredentials, "invalid credentials")
		}

		return nil, errors.Wrap(err, "checking email")
	}

	return &res, nil
}
