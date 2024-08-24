package handler

import (
	"context"

	"github.com/abdinep/rent-a-car-grpc/car-service/internal/service"
	cars "github.com/abdinep/rent-a-car-grpc/car-service/proto"
)

type CarHandler struct {
	cars.UnimplementedCarServiceServer
	Service service.CarService
}

func NewCarHandler(svc service.CarService) *CarHandler {
	return &CarHandler{Service: svc}
}
func (h *CarHandler) AddCar(ctx context.Context, req *cars.AddCarRequest) (*cars.AddCarResponse, error) {
	Id, err := h.Service.AddCar(req.Model, req.Make, req.Year)
	if err != nil {
		return nil, err
	}
	return &cars.AddCarResponse{Id: Id}, nil
}
func (h *CarHandler) ListCars(ctx context.Context, req *cars.ListCarsRequest) (*cars.ListCarsResponse, error) {
	car, err := h.Service.ListCars()
	if err != nil {
		return nil, err
	}
	var carsResponse []*cars.Car
	for _, car := range *car {
		carsResponse = append(carsResponse, &cars.Car{
			Id:    car.ID,
			Model: car.Model,
			Make:  car.Make,
			Year:  car.Year,
		})
	}
	return &cars.ListCarsResponse{Cars: carsResponse}, nil
}
