package service

import (
	"context"

	"github.com/eugenshima/moviori/internal/model"
)

type AuthService struct {
	srv AuthRepositoryInterface
}

func NewAuthService(srv AuthRepositoryInterface) *AuthService {
	return &AuthService{srv: srv}
}

type AuthRepositoryInterface interface {
	InsertUser(ctx context.Context, auth *model.AuthModel) error
	GetUserByID(ctx context.Context) error
}

func (s *AuthService) LoginService(ctx context.Context, auth *model.AuthModel) error {
	return s.srv.InsertUser(ctx, auth)
}
