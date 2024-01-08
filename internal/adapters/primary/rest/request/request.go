package request

import (
	"regexp"
	"unicode/utf8"

	"cars/internal/domain/auth"
	"cars/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
)

var emailRegexp = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
var passwordRegexp = regexp.MustCompile(`[.,\(\);:\\\/\[\]\{\}@$!%*#?&=]`)

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
	if !emailRegexp.MatchString(r.Email) {
		return errors.Wrap(errors.InvalidRequestContent, "email")
	}

	if !passwordRegexp.MatchString(r.Password) || utf8.RuneCountInString(r.Password) < 6 {
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
