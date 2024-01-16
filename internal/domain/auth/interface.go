package auth

// Usecase

type IUsecase interface {
	Signup(*SignupReq) (*SignupRes, error)
	Signin(*SigninReq) (*SigninRes, error)
	Refresh(*RefreshReq) (*RefreshRes, error)
	Signout(*SignoutReq) (*SignoutRes, error)
	SignoutAll(*SignoutAllReq) (*SignoutAllRes, error)
}

// Adapters

type IRepository interface {
	Signup(*RepoSignupReq) (*RepoSignupRes, error)
	Signin(*RepoSigninReq) (*RepoSigninRes, error)
}

type IToken interface {
	Create(id uint64) (token string, err error)
	Parse(token string) (claims Claims, err error)
	ParseExpired(token string) (claims Claims, err error)
	SaveRefresh(token string, id uint64) error
	RemoveRefresh(token string, id uint64) error
	RemoveAllRefresh(id uint64) error
}
