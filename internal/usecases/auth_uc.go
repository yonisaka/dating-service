package usecases

import (
	"context"

	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/jwtx"
)

// AuthUsecase is an interface for auth usecase
type AuthUsecase interface {
	Login(ctx context.Context, req presentations.LoginRequest) (*presentations.LoginResponse, error)
	Register(ctx context.Context, req presentations.RegisterRequest) (*presentations.RegisterResponse, error)
}

func NewAuthUsecase(
	userRepo repository.UserRepo,
) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
	}
}

type authUsecase struct {
	userRepo repository.UserRepo
	token    jwtx.JWTImpl
}
