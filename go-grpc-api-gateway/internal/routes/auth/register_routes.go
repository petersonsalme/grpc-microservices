package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/internal/clients"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/internal/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *clients.AuthClient {
	svc := &clients.AuthClient{
		Client: clients.InitAuthClient(c),
	}

	routes := r.Group("/auth")

	routes.POST("/register", func(ctx *gin.Context) {
		Register(ctx, svc.Client)
	})

	routes.POST("/login", func(ctx *gin.Context) {
		Login(ctx, svc.Client)
	})

	return svc
}
