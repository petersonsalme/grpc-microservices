package order

import (
	"github.com/gin-gonic/gin"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/internal/clients"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/internal/config"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/internal/middleware"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, a *clients.AuthClient) {
	m := middleware.InitAuthMiddleware(a)
	svc := &clients.OrderClient{
		Client: clients.InitOrderClient(c),
	}

	routes := r.Group("/orders")
	routes.Use(m.AuthRequired)

	routes.POST("/", func(ctx *gin.Context) {
		Create(ctx, svc.Client)
	})
}
