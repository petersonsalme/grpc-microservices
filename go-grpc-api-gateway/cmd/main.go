package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/internal/config"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/internal/routes"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at load config", err)
	}

	r := gin.Default()

	routes.RegisterAll(r, &c)

	r.Run(c.Port)
}
