package service

import (
	"context"
	"grpc-starter/internal/model"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (AuthService) Login(ctx context.Context, username string, password string) (*model.UserModel, error) {

	return &model.UserModel{Id: 1}, nil
}
