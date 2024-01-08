package cars

// Usecase

type IUsecase interface {
	Create(*CreateReq) (*CreateRes, error)
	Get(*GetReq) (*GetRes, error)
	UpdateColor(*UpdateColorReq) (*UpdateColorRes, error)
	Delete(*DeleteReq) (*DeleteRes, error)
}

// Adapters

type IRepository interface {
	Create(*RepoCreateReq) (*RepoCreateRes, error)
	Get(*RepoGetReq) (*RepoGetRes, error)
	UpdateColor(*RepoUpdateColorReq) (*RepoUpdateColorRes, error)
	Delete(*RepoDeleteReq) (*RepoDeleteRes, error)
}
