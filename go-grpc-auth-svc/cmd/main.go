package main

import (
	"fmt"
	"log"
	"net"

	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/pkg/pb"
	"github.com/petersonsalme/grpc-microservices/go-grpc-auth-svc/internal/config"
	"github.com/petersonsalme/grpc-microservices/go-grpc-auth-svc/internal/db"
	"github.com/petersonsalme/grpc-microservices/go-grpc-auth-svc/internal/jwt"
	"github.com/petersonsalme/grpc-microservices/go-grpc-auth-svc/internal/server"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	jwt := jwt.Wrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 1,
	}

	listener, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Service on", c.Port)

	s := server.Server{
		H:   h,
		JWT: jwt,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
