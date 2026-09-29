package order

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/pkg/pb"
)

type CreateRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

func Create(ctx *gin.Context, c pb.OrderServiceClient) {
	var body CreateRequest
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, _ := ctx.Get("userId")
	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: body.ProductID,
		Quantity:  body.Quantity,
		UserId:    userID.(int64),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
