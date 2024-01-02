package carsDomain

// Usecase

type IUsecase interface {
	Create(*UcaseReqCreate) error
	Get(*UcaseReqGet) (*UcaseResGet, error)
	UpdateColor(*UcaseReqUpdateColor) error
	Delete(*UcaseReqDelete) error
}

// Repository

type IRepository interface {
	Create(*RepoReqCreate) error
	Get(*RepoReqGet) (*RepoResGet, error)
	UpdateColor(*RepoReqUpdateColor) error
	Delete(*RepoReqDelete) error
}
