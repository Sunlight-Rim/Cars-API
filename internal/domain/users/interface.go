package users

import "cars/internal/domain/auth"

// Usecase

type IUsecase interface {
	GetMe(*GetMeReq) (*GetMeRes, error)
	UpdateInfo(*UpdateInfoReq) (*UpdateInfoRes, error)
	UpdatePassword(*UpdatePasswordReq) error
	Delete(*DeleteReq) (*DeleteRes, error)
}

// Adapters

type IRepository interface {
	GetMe(*RepoGetMeReq) (*RepoGetMeRes, error)
	UpdateInfo(*RepoUpdateInfoReq) (*RepoUpdateInfoRes, error)
	UpdatePassword(*RepoUpdatePasswordReq) error
	Delete(*RepoDeleteReq) (*RepoDeleteRes, error)
}

type IToken interface {
	Parse(token string) (claims *auth.Claims, err error)
}
