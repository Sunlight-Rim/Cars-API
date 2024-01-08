package carsRepository

import (
	"cars/internal/domain/cars"
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

func (r *repository) Create(req *cars.RepoCreateReq) (*cars.RepoCreateRes, error) {
	return nil, nil
}
func (r *repository) Get(req *cars.RepoGetReq) (*cars.RepoGetRes, error) {
	return nil, nil
}
func (r *repository) UpdateColor(req *cars.RepoUpdateColorReq) (*cars.RepoUpdateColorRes, error) {
	return nil, nil
}
func (r *repository) Delete(req *cars.RepoDeleteReq) (*cars.RepoDeleteRes, error) {
	return nil, nil
}
