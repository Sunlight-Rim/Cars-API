package carsRepository

import (
	carsEntity "cars/internal/entities/cars"
	"database/sql"
)

type repository struct {
	psql *sql.DB
}

func New(postgres *sql.DB) *repository {
	return &repository{
		psql: postgres,
	}
}

func (r *repository) Create(req *carsEntity.RepositoryReqCreate) (*carsEntity.RepositoryResCreate, error) {
	return nil, nil
}
func (r *repository) Get(req *carsEntity.RepositoryReqGet) (*carsEntity.RepositoryResGet, error) {
	return nil, nil
}
func (r *repository) UpdateColor(req *carsEntity.RepositoryReqUpdateColor) (*carsEntity.RepositoryResUpdateColor, error) {
	return nil, nil
}
func (r *repository) Delete(req *carsEntity.RepositoryReqDelete) (*carsEntity.RepositoryResDelete, error) {
	return nil, nil
}
