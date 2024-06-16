package service

import (
	"context"
	"fmt"

	"github.com/eugenshima/moviori/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	srv AuthRepositoryInterface
}

func NewAuthService(srv AuthRepositoryInterface) *AuthService {
	return &AuthService{srv: srv}
}

type AuthRepositoryInterface interface {
	InsertNewUser(ctx context.Context, auth *model.HashedLogin) error
	GetUserByID(ctx context.Context) error
	GetUserByLogin(ctx context.Context, login string) (*model.FullUserModel, error)
}

func (s *AuthService) LoginService(ctx context.Context, login *model.UserModel) (*model.FullUserModel, error) {
	user, err := s.srv.GetUserByLogin(ctx, login.Login)
	if err != nil {
		return nil, fmt.Errorf("GetUserByLogin: %w", err)
	}
	isRight := CheckPasswordHash(login.Password, user.Password)
	if !isRight {
		return nil, fmt.Errorf("CheckPasswordHash: wrong password")
	}
	return user, nil
}

func (s *AuthService) SignupService(ctx context.Context, signup *model.UserModel) error {
	hash, err := HashPassword(signup.Password)
	if err != nil {
		return fmt.Errorf("SignupService: %w", err)
	}
	user := &model.HashedLogin{
		Login:    signup.Login,
		Password: hash,
	}
	return s.srv.InsertNewUser(ctx, user)
}

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
