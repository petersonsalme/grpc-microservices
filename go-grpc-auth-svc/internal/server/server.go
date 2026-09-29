package server

import (
	"context"
	"net/http"

	"github.com/petersonsalme/grpc-microservices/go-grpc-api-gateway/pkg/pb"

	"github.com/petersonsalme/grpc-microservices/go-grpc-auth-svc/internal/db"
	"github.com/petersonsalme/grpc-microservices/go-grpc-auth-svc/internal/encrypt"
	"github.com/petersonsalme/grpc-microservices/go-grpc-auth-svc/internal/jwt"
	"github.com/petersonsalme/grpc-microservices/go-grpc-auth-svc/internal/models"
)

type Server struct {
	H   db.Handler
	JWT jwt.Wrapper
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&models.User{}); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	s.H.DB.Create(&models.User{
		Email:    req.Email,
		Password: encrypt.Password(req.Password),
	})

	return &pb.RegisterResponse{Status: http.StatusCreated}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User
	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	if !encrypt.IsEquals(req.Password, user.Password) {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token, _ := s.JWT.GenerateToken(user)
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.JWT.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User
	if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.ID,
	}, nil
}
