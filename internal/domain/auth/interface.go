package auth

// Usecase

type IUsecase interface {
	Signup(*SignupReq) (*SignupRes, error)
	Signin(*SigninReq) (*SigninRes, error)
	Signout(*SignoutReq) (*SignoutRes, error)
	Refresh(*RefreshReq) (*RefreshRes, error)
}

// Adapters

type IRepository interface {
	Signup(*RepoSignupReq) (*RepoSignupRes, error)
	Signin(*RepoSigninReq) (*RepoSigninRes, error)
}

type IToken interface {
	Create(claims map[string]string) (token string, err error)
	Parse(token string) (claims map[string]string, err error)
	ParseExpired(token string) (claims map[string]string, err error)
	Save(token string) error
	Delete(token string) error
}
