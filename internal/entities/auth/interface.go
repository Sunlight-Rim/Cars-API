package authEntity

// Usecase

type IUsecase interface {
	Signup(*UsecaseReqSignup) (*UsecaseResSignup, error)
	Signin(*UsecaseReqSignin) (*UsecaseResSignin, error)
	Signout(*UsecaseReqSignout) (*UsecaseResSignout, error)
	Refresh(*UsecaseReqRefresh) (*UsecaseResRefresh, error)
}

// Adapters

type IRepository interface {
	Signup(*RepositoryReqSignup) (*RepositoryResSignup, error)
	Signin(*RepositoryReqSignin) (*RepositoryResSignin, error)
}

type IToken interface {
	Create(claims map[string]string) (token string, err error)
	Parse(token string) (claims map[string]string, err error)
	ParseExpired(token string) (claims map[string]string, err error)
	Save(token string) error
	Delete(token string) error
}
