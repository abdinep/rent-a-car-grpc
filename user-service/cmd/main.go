package main

import (
	"log"
	"net"

	"github.com/abdinep/rent-a-car-grpc/user-service/internal/handler"
	"github.com/abdinep/rent-a-car-grpc/user-service/internal/repository"
	"github.com/abdinep/rent-a-car-grpc/user-service/internal/service"
	user "github.com/abdinep/rent-a-car-grpc/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const carServiceAddress = "localhost:50052"

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	connectCarService, err := grpc.Dial(carServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to car service: %v", err)
	}
	defer connectCarService.Close()

	repo := repository.NewUserRepository()
	userService := service.NewUserService(repo, connectCarService)
	userHandler := handler.NewUserHandler(userService)

	user.RegisterUserServiceServer(s, userHandler)
	reflection.Register(s)

	log.Printf("Server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
