package usersEntity

// Usecase

type IUsecase interface {
	Get(*UsecaseReqGet) (*UsecaseResGet, error)
	UpdateInfo(*UsecaseReqUpdateInfo) (*UsecaseResUpdateInfo, error)
	UpdatePassword(*UsecaseReqUpdatePassword) error
	Delete(*UsecaseReqDelete) (*UsecaseResDelete, error)
}

// Adapters

type IRepository interface {
	Get(*RepositoryReqGet) (*RepositoryResGet, error)
	UpdateInfo(*RepositoryReqUpdateInfo) (*RepositoryResUpdateInfo, error)
	UpdatePassword(*RepositoryReqUpdatePassword) error
	Delete(*RepositoryReqDelete) (*RepositoryResDelete, error)
}
