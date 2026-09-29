package product

import (
	"github.com/gin-gonic/gin"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/internal/clients"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/internal/config"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/internal/middleware"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, a *clients.AuthClient) {
	m := middleware.InitAuthMiddleware(a)

	svc := &clients.ProductClient{
		Client: clients.InitProductClient(c),
	}

	routes := r.Group("/products")
	routes.Use(m.AuthRequired)

	routes.POST("/", func(ctx *gin.Context) {
		Create(ctx, svc.Client)
	})

	routes.GET("/:id", func(ctx *gin.Context) {
		FindOne(ctx, svc.Client)
	})
}
