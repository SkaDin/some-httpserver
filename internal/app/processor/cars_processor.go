package processor

import (
	"errors"
	"some-httpserver/internal/app/db"
	"some-httpserver/internal/app/models"
)

type CarsProcessor struct {
	storage *db.CarsStorage
}

func NewCarsProcessor(storage *db.CarsStorage) *CarsProcessor {
	processor := &CarsProcessor{
		storage: storage,
	}
	return processor
}

func (processor *CarsProcessor) CreateCar(car models.Car) error {
	if car.Colour == "" {
		return errors.New("colour should not be empty")
	}
	if car.Brand == "" {
		return errors.New("brand should not be empty")
	}
	if car.LicencePlate == "" {
		return errors.New("licence plate should not be empty")
	}
	if car.Owner.Id <= 0 {
		return errors.New("owner id shall be filled")
	}
	return processor.storage.CreateCar(car)
}

func (processor *CarsProcessor) FindCar(id uint64) (models.Car, error) {
	car := processor.storage.GetCarById(id)

	if car.Id != id {
		return car, errors.New("car not found")
	}
	return car, nil
}

func (processor *CarsProcessor) ListCars(
	userId uint64,
	brandFilter string,
	colourFilter string,
	licenceFilter string,
) ([]models.Car, error) {
	return processor.storage.GetCarsList(userId, brandFilter, colourFilter, licenceFilter), nil
}
