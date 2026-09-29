package clients

import (
	"context"
	"fmt"

	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/pkg/pb"
	"google.golang.org/grpc"
)

type ProductClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return ProductClient{
		Client: pb.NewProductServiceClient(cc),
	}
}

func (c *ProductClient) FindOne(productId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productId,
	}

	return c.Client.FindOne(context.Background(), req)
}

func (c *ProductClient) DecreaseStock(productId int64, orderId int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	}

	return c.Client.DecreaseStock(context.Background(), req)
}
