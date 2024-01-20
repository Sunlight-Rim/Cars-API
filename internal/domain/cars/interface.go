package cars

import "cars/internal/domain/auth"

// Usecase

type IUsecase interface {
	Create(*CreateReq) (*CreateRes, error)
	Get(*GetReq) (*GetRes, error)
	Update(*UpdateReq) (*UpdateRes, error)
	Delete(*DeleteReq) (*DeleteRes, error)
}

// Adapters

type IRepository interface {
	Create(*RepoCreateReq) (*RepoCreateRes, error)
	Get(*RepoGetReq) (*RepoGetRes, error)
	Update(*RepoUpdateReq) (*RepoUpdateRes, error)
	Delete(*RepoDeleteReq) (*RepoDeleteRes, error)
}

type IToken interface {
	Parse(token string) (claims *auth.Claims, err error)
}
