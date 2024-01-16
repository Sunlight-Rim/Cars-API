package request

import (
	"regexp"
	"unicode/utf8"

	"cars/internal/domain/auth"
	"cars/pkg/errors"
)

var emailRegex = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
var passwordRegex = regexp.MustCompile(`[.,\(\);:\\\/\[\]\{\}@$!%*#?&=]`)

// Signup

type Signup struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

func (r *Signup) Validate() error {
	if !emailRegex.MatchString(r.Email) {
		return errors.Wrap(errors.InvalidRequestContent, "email")
	}

	if !passwordRegex.MatchString(r.Password) || utf8.RuneCountInString(r.Password) < 6 {
		return errors.Wrap(errors.InvalidRequestContent, "password")
	}

	return nil
}

func (r *Signup) ToEntity() *auth.SignupReq {
	return &auth.SignupReq{
		Username: r.Username,
		Email:    r.Email,
		Address:  r.Address,
		Password: r.Password,
	}
}

// Signin

type Signin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *Signin) ToEntity() *auth.SigninReq {
	return &auth.SigninReq{
		Email:    r.Email,
		Password: r.Password,
	}
}
