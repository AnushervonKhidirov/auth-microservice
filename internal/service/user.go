package service

import (
	"auth/internal/repository"
	"context"
)

type UserService struct {
	userRepository  *repository.UserRepository
	tokenRepository *repository.TokenRepository
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s UserService) Get(ctx context.Context) {
}

func (s UserService) GetAll(ctx context.Context) {
}

func (s UserService) Create(ctx context.Context) {
}

func (s UserService) Update(ctx context.Context) {
}

func (s UserService) Delete(ctx context.Context) {
}
