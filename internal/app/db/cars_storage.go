package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"some-httpserver/internal/app/models"
)

type CarsStorage struct {
	databasePool *pgxpool.Pool
}
type userCar struct {
	UserId       uint64 `db:"userid"`
	Name         string
	Rank         string
	CarId        uint64 `db:"carid"`
	Brand        string
	Colour       string
	LicencePlate string
}

func converterJoinedQueryToCar(input userCar) models.Car {
	return models.Car{
		Id:           input.CarId,
		Colour:       input.Colour,
		Brand:        input.Brand,
		LicencePlate: input.LicencePlate,
		Owner: models.User{
			Id:   input.UserId,
			Name: input.Name,
			Rank: input.Rank,
		},
	}
}

func NewCarsStorage(pool *pgxpool.Pool) *CarsStorage {
	storage := &CarsStorage{
		databasePool: pool,
	}
	return storage
}

func (storage *CarsStorage) GetCarsList(
	userIdFilter uint64,
	brandFilter string,
	colourFilter string,
	licenceFilter string,
) []models.Car {
	query := `SELECT
				users.id AS userid,
				users.name,
				users.rank,
				c.id AS carid,
				c.brand,
				c.colour,
				c.licence_plate 
			FROM users 
			JOIN cars c on users.id = c.user_id
			WHERE 1 = 1`

	placeholderNum := 1
	args := make([]any, 0)
	if userIdFilter != 0 {
		query += fmt.Sprintf(" AND users.id = $%d", placeholderNum)
		args = append(args, userIdFilter)
		placeholderNum++
	}
	if brandFilter != "" {
		query += fmt.Sprintf(" AND brand ILIKE $%d", placeholderNum)
		args = append(args, fmt.Sprintf("%%%s%%", brandFilter))
		placeholderNum++
	}
	if colourFilter != "" {
		query += fmt.Sprintf(" AND colour ILIKE $%d", placeholderNum)
		args = append(args, fmt.Sprintf("%%%s%%", colourFilter))
		placeholderNum++
	}
	if licenceFilter != "" {
		query += fmt.Sprintf(" AND licence_plate ILIKE $%d", placeholderNum)
		args = append(args, fmt.Sprintf("%%%s%%"), licenceFilter)
		placeholderNum++
	}

	var dbResult []userCar

	err := pgxscan.Select(context.Background(), storage.databasePool, &dbResult, query, args...)

	if err != nil {
		log.Errorln(err)
	}
	result := make([]models.Car, len(dbResult))

	for ind, dbEntity := range dbResult {
		result[ind] = converterJoinedQueryToCar(dbEntity)
	}
	return result
}

func (storage *CarsStorage) GetCarById(id uint64) models.Car {
	query := `SELECT
				users.id AS userid,
				users.name,
				users.rank,
				c.id AS carid,
				c.brand,
				c.colour,
				c.licence_plate
			  FROM users
				JOIN cars c 
					ON users.id = c.user_id
			  WHERE c.id = $1`

	var result userCar

	err := pgxscan.Get(context.Background(), storage.databasePool, &result, query, id)

	if err != nil {
		log.Errorln(err)
	}
	return converterJoinedQueryToCar(result)
}

func (storage *CarsStorage) CreateCar(car models.Car) error {
	ctx := context.Background()
	tx, err := storage.databasePool.Begin(ctx)
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Errorln(err)
		}
	}()
	query := `SELECT id 
    		  FROM users
    		  WHERE id = $1`
	id := -1

	err = pgxscan.Get(ctx, tx, &id, query, car.Owner.Id)

	if err != nil {
		log.Errorln(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Errorln(err)
		}
		return err
	}

	if id == -1 {
		err = errors.New("user not found")
	}
	insertQuery := `INSERT INTO cars(user_id, colour, brand, licence_plate)
						VALUES ($1,$2,$3,$4)`
	_, err = tx.Exec(context.Background(), insertQuery, car.Owner.Id, car.Colour, car.Brand, car.LicencePlate)

	if err != nil {
		log.Errorln(err)
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Errorln(err)
		}
		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Errorln(err)
	}
	return err
}
