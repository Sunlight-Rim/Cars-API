package authRepository

import (
	"database/sql"

	"cars/internal/domain/auth"
	"cars/pkg/errors"

	"github.com/lib/pq"
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
	var res auth.RepoSignupRes

	// Add user
	if err := r.psql.QueryRow(`
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
		// Unique constraint violation
		if pqError := new(pq.Error); errors.As(err, &pqError) && pqError.Code == "23505" {
			return nil, errors.Wrap(errors.EmailIsBusy, "busy email")
		}

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
			password_hash = $2
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
