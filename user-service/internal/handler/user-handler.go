package handler

import (
	"context"

	"github.com/abdinep/rent-a-car-grpc/user-service/internal/service"
	user "github.com/abdinep/rent-a-car-grpc/user-service/proto"
)

type UserHandler struct {
	user.UnimplementedUserServiceServer
	Service service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{Service: svc}
}
func (h *UserHandler) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	err := h.Service.Register(req.Username, req.Password, req.Email)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResponse{Message: "User registered successfully"}, nil
}
func (h *UserHandler) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	token, err := h.Service.Login(req.Password, req.Username)
	if err != nil {
		return nil, err
	}
	return &user.LoginResponse{Token: token}, nil
}
func (h *UserHandler) Getcars(ctx context.Context, req *user.CarRequest) (*user.CarResponse, error) {
	res, err := h.Service.Getcars()
	if err != nil {
		return nil, err
	}
	var carlist []*user.Car
	for _, list := range res {
		carlist = append(carlist, &user.Car{
			Id:    list.Id,
			Make:  list.Make,
			Model: list.Model,
			Year:  list.Year,
		})

	}
	return &user.CarResponse{Cars: carlist}, nil
}
