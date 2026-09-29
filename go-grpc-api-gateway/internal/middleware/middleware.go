package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/internal/clients"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/pkg/pb"
)

type AuthMiddlewareConfig struct {
	svc *clients.AuthClient
}

func InitAuthMiddleware(svc *clients.AuthClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		fmt.Println("Unauthorized")
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}
