package request

import (
	"regexp"
	"unicode/utf8"

	"cars/internal/domain/auth"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
	easyjson "github.com/mailru/easyjson"
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

func NewSignup(c echo.Context) (*Signup, error) {
	var r Signup

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	if err := r.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation")
	}

	return &r, nil
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

func NewSignin(c echo.Context) (*Signin, error) {
	var r Signin

	if err := easyjson.UnmarshalFromReader(c.Request().Body, &r); err != nil {
		return nil, errors.Wrapf(errors.InvalidRequestFormat, "parsing, %v", err)
	}

	return &r, nil
}

func (r *Signin) ToEntity() *auth.SigninReq {
	return &auth.SigninReq{
		Email:    r.Email,
		Password: r.Password,
	}
}

// Refresh

// Signout

// Singout all
