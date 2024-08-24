package repository

import (
	"github.com/abdinep/rent-a-car-grpc/car-service/internal/model"
	"github.com/google/uuid"
)

type CarRepository interface {
	AddCars(models, make string, year int32) (string, error)
	ListingCars() (*[]model.Cars, error)
}
type carRepository struct {
	cars map[string]*model.Cars
}

func NewCarRepository() CarRepository {
	return &carRepository{cars: make(map[string]*model.Cars)}
}
func (r *carRepository) AddCars(models, make string, year int32) (string, error) {
	carId := uuid.New().String()
	car := &model.Cars{
		ID:    carId,
		Model: models,
		Make:  make,
		Year:  year,
	}
	r.cars[carId] = car
	return carId, nil
}
func (r *carRepository) ListingCars() (*[]model.Cars, error) {
	var carlist []model.Cars
	for _, car := range r.cars {
		carlist = append(carlist, *car)
	}
	return &carlist, nil
}
