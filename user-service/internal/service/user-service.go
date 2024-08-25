package service

import (
	"context"
	"errors"

	cars "github.com/abdinep/rent-a-car-grpc/user-service/proto/client"
	"google.golang.org/grpc"
)

type User struct {
	ID       string
	Username string
	Password string
	Email    string
}

type UserRepository interface {
	CreateUser(username, password, email string) error
	FindUserByUsername(username string) (*User, error)
}

type UserService interface {
	Register(username, password, email string) error
	Login(username, password string) (string, error)
	Getcars() ([]*cars.Car, error)
}

type userService struct {
	repo      UserRepository
	carClient cars.CarServiceClient
}

func NewUserService(repo UserRepository, carclient *grpc.ClientConn) UserService {
	carsclient := cars.NewCarServiceClient(carclient)
	return &userService{
		repo:      repo,
		carClient: carsclient,
	}
}
func (s *userService) Register(username, password, email string) error {
	err := s.repo.CreateUser(username, password, email)
	if err != nil {
		return err
	}
	return nil
}
func (s *userService) Login(password, username string) (string, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil || user.Password != password {
		return "", errors.New("invalid username or password")
	}
	token := "token-" + user.ID
	return token, nil
}
func (s *userService) Getcars() ([]*cars.Car, error) {
	res,err := s.carClient.ListCars(context.Background(),&cars.ListCarsRequest{})
	if err != nil{
		return nil,err
	}
	return res.Cars,nil
}
