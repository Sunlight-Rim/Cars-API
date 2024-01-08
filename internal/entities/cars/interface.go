package carsEntity

// Usecase

type IUsecase interface {
	Create(*UsecaseReqCreate) (*UsecaseResCreate, error)
	Get(*UsecaseReqGet) (*UsecaseResGet, error)
	UpdateColor(*UsecaseReqUpdateColor) (*UsecaseResUpdateColor, error)
	Delete(*UsecaseReqDelete) (*UsecaseResDelete, error)
}

// Adapters

type IRepository interface {
	Create(*RepositoryReqCreate) (*RepositoryResCreate, error)
	Get(*RepositoryReqGet) (*RepositoryResGet, error)
	UpdateColor(*RepositoryReqUpdateColor) (*RepositoryResUpdateColor, error)
	Delete(*RepositoryReqDelete) (*RepositoryResDelete, error)
}
