package carsRepository

import (
	carsDomain "cars/internal/domain/cars"
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

func (r *repository) Create(req *carsDomain.RepoReqCreate) error                     { return nil }
func (r *repository) Get(req *carsDomain.RepoReqGet) (*carsDomain.RepoResGet, error) { return nil, nil }
func (r *repository) UpdateColor(req *carsDomain.RepoReqUpdateColor) error           { return nil }
func (r *repository) Delete(req *carsDomain.RepoReqDelete) error                     { return nil }
