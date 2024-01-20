package cars

// Usecase
type IUsecase interface {
	Create(*CreateReq) (*CreateRes, error)
	Get(*GetReq) (*GetRes, error)
	Update(*UpdateReq) (*UpdateRes, error)
	Delete(*DeleteReq) (*DeleteRes, error)
}

// Repository
type IRepository interface {
	Create(*RepoCreateReq) (*RepoCreateRes, error)
	Get(*RepoGetReq) (*RepoGetRes, error)
	Update(*RepoUpdateReq) (*RepoUpdateRes, error)
	Delete(*RepoDeleteReq) (*RepoDeleteRes, error)
}
