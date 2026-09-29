package server

import (
	"context"
	"net/http"

	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/pkg/pb"
	"github.com/petersonsalme/grpc-microservices/go-grpc-order-svc/internal/clients"
	"github.com/petersonsalme/grpc-microservices/go-grpc-order-svc/internal/db"
	"github.com/petersonsalme/grpc-microservices/go-grpc-order-svc/internal/models"
)

type Server struct {
	H              db.Handler
	ProductService clients.ProductClient
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	product, err := s.ProductService.FindOne(req.ProductId)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if product.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{
			Status: product.Status,
			Error:  product.Error,
		}, nil
	} else if product.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  "Stock too less",
		}, nil
	}

	order := models.Order{
		Price:     product.Data.Price,
		ProductID: product.Data.Id,
		UserID:    req.UserId,
	}
	s.H.DB.Create(&order)

	res, err := s.ProductService.DecreaseStock(req.ProductId, order.ID)
	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if res.Status == http.StatusConflict {
		s.H.DB.Delete(&models.Order{}, order.ID)
		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  res.Error,
		}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.ID,
	}, nil
}
