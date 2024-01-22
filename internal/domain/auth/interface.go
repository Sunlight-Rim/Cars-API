package auth

// Usecase
type IUsecase interface {
	Signup(*SignupReq) (*SignupRes, error)
	Signin(*SigninReq) (*SigninRes, error)
	Refresh(*RefreshReq) (*RefreshRes, error)
	Sessions(*SessionsReq) (*SessionsRes, error)
	Signout(*SignoutReq) (*SignoutRes, error)
	SignoutAll(*SignoutAllReq) (*SignoutAllRes, error)
}

// Repository
type IRepository interface {
	Signup(*RepoSignupReq) (*RepoSignupRes, error)
	Signin(*RepoSigninReq) (*RepoSigninRes, error)
	IsDeleted(*RepoIsDeletedReq) (*RepoIsDeletedRes, error)
}

// Token service
type IToken interface {
	Create(userID uint64) (token string, err error)
	Parse(token string) (claims *Claims, err error)
	ParseExpired(token string) (claims *Claims, err error)
	StoreUserRefresh(userID uint64, token string) error
	ListUserRefresh(userID uint64) (tokens []string, err error)
	RevokeUserRefresh(userID uint64, token string) error
	RevokeUserRefreshAll(userID uint64) (tokens []string, err error)
}
