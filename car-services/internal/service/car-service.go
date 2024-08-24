package service

import (
	"github.com/abdinep/rent-a-car-grpc/car-service/internal/model"
	"github.com/abdinep/rent-a-car-grpc/car-service/internal/repository"
)

type CarService interface {
	AddCar(model, make string, year int32) (string, error)
	ListCars() (*[]model.Cars, error)
}
type carService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{repo: repo}
}
func (s *carService) AddCar(model, make string, year int32) (string, error) {
	car, err := s.repo.AddCars(model, make, year)
	if err != nil {
		return "", err
	}
	return car, nil
}
func (s *carService) ListCars() (*[]model.Cars, error) {
	cars, err := s.repo.ListingCars()
	if err != nil {
		return nil, err
	}
	return cars, nil
}
