package main

import (
	"log"
	"net"

	"github.com/abdinep/rent-a-car-grpc/car-service/internal/handler"
	"github.com/abdinep/rent-a-car-grpc/car-service/internal/repository"
	"github.com/abdinep/rent-a-car-grpc/car-service/internal/service"
	cars "github.com/abdinep/rent-a-car-grpc/car-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	repo := repository.NewCarRepository()
	carservice := service.NewCarService(repo)
	carHandler := handler.NewCarHandler(carservice)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	cars.RegisterCarServiceServer(s, carHandler)

	reflection.Register(s)

	log.Printf("Server listening on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
