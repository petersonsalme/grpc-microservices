package main

import (
	"fmt"
	"log"
	"net"

	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/pkg/pb"
	"github.com/petersonsalme/grpc-microservices/go-grpc-product-svc/internal/config"
	"github.com/petersonsalme/grpc-microservices/go-grpc-product-svc/internal/db"
	"github.com/petersonsalme/grpc-microservices/go-grpc-product-svc/internal/server"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	s := server.Server{H: db.Init(c.DBUrl)}

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", c.Port)

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
