package service

import (
	"auth/internal/repository"
	"context"
)

type AuthService struct {
	userRepository  *repository.UserRepository
	tokenRepository *repository.TokenRepository
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s AuthService) SignUp(ctx context.Context) {
}

func (s AuthService) SignIn(ctx context.Context) {
}

func (s AuthService) SignOut(ctx context.Context) {
}
