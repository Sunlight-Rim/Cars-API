package usersDomain

// Usecase

type IUsecase interface {
	Register(*UcaseReqRegister) error
	Login(*UcaseReqLogin) (*UcaseResLogin, error)
	Logout(*UcaseReqLogout) error
	Refresh(*UcaseReqRefresh) (*UcaseResRefresh, error)

	Get(*UcaseReqGet) (*UcaseResGet, error)
	UpdateInfo(*UcaseReqUpdateInfo) error
	UpdatePassword(*UcaseReqUpdatePassword) error
	Delete(*UcaseReqDelete) error
}

// Repository

type IRepository interface {
	Register(*RepoReqRegister) error
	Login(*RepoReqLogin) (*RepoResLogin, error)
	Logout(*RepoReqLogout) error
	Refresh(*RepoReqRefresh) (*RepoResRefresh, error)

	Get(*RepoReqGet) (*RepoResGet, error)
	UpdateInfo(*RepoReqUpdateInfo) error
	UpdatePassword(*RepoReqUpdatePassword) error
	Delete(*RepoReqDelete) error
}

// Token

type IToken interface {
	Create(claims map[string]string) (token string, err error)
	Parse(token string) (claims map[string]string, err error)
	ParseExpired(token string) (claims map[string]string, err error)
	Save(token string) error
	Delete(token string) error
}
