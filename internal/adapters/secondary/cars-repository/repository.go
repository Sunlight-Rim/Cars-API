package carsRepository

import (
	"database/sql"

	"cars/internal/domain/cars"
	"cars/pkg/errors"
)

type repository struct {
	psql *sql.DB
}

func New(postgres *sql.DB) *repository {
	return &repository{
		psql: postgres,
	}
}

func (r *repository) Create(req *cars.RepoCreateReq) (_ *cars.RepoCreateRes, err error) {
	var res cars.RepoCreateRes

	// Start transacton
	tx, err := r.psql.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "begin transaction")
	}

	// End transacton
	defer func() {
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				err = errors.Wrapf(errRollback, "rollback, %v", err)
			}
		} else {
			if errCommit := tx.Commit(); errCommit != nil {
				err = errors.Wrap(errCommit, "commit")
			}
		}
	}()

	// Check if plate is busy
	var plateBusy bool

	if err := tx.QueryRow(`
		SELECT true
		FROM api.cars
		WHERE plate = $1
	`, req.Plate,
	).Scan(&plateBusy); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, "checking plate")
	}

	if plateBusy {
		return nil, errors.Wrap(errors.PlateIsBusy, "busy plate")
	}

	// Create new car
	if err := r.psql.QueryRow(`
		INSERT INTO api.cars(
			user_id,
			plate,
			model,
			color
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		req.UserID,
		req.Plate,
		req.Model,
		req.Color,
	).Scan(&res.ID); err != nil {
		return nil, errors.Wrap(err, "adding car")
	}

	return &res, nil
}

func (r *repository) Get(req *cars.RepoGetReq) (*cars.RepoGetRes, error) {
	var res cars.RepoGetRes

	rows, err := r.psql.Query(`
		SELECT
			id,
			plate,
			model,
			color
		FROM api.cars
		WHERE user_id = $1
	`, req.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "getting cars")
	}
	defer rows.Close()

	for rows.Next() {
		var car cars.Car

		if err := rows.Scan(
			&car.ID,
			&car.Plate,
			&car.Model,
			&car.Color,
		); err != nil {
			return nil, errors.Wrap(err, "scan car")
		}

		res.Cars = append(res.Cars, &car)
	}

	return &res, nil
}

func (r *repository) Update(req *cars.RepoUpdateReq) (*cars.RepoUpdateRes, error) {
	var res = cars.RepoUpdateRes{Car: cars.Car{
		ID:    req.CarID,
		Model: req.Model,
		Color: req.Color,
	}}

	if err := r.psql.QueryRow(`
		UPDATE api.cars
		SET
			model = $1,
			color = $2
		WHERE
			id = $3 AND
			user_id = $4
		RETURNING plate
	`,
		req.Model,
		req.Color,
		req.CarID,
		req.UserID,
	).Scan(&res.Plate); err != nil {
		return nil, errors.Wrap(err, "updating car")
	}

	return &res, nil
}

func (r *repository) Delete(req *cars.RepoDeleteReq) (*cars.RepoDeleteRes, error) {
	var res = cars.RepoDeleteRes{Car: cars.Car{
		ID: req.CarID,
	}}

	if err := r.psql.QueryRow(`
		DELETE FROM api.cars
		WHERE
			id = $1 AND
			user_id = $2
		RETURNING
			plate,
			model,
			color
	`,
		req.CarID,
		req.UserID,
	).Scan(
		&res.Plate,
		&res.Model,
		&res.Color,
	); err != nil {
		return nil, errors.Wrap(err, "deleting car")
	}

	return &res, nil
}
