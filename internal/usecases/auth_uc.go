package usecases

import (
	"context"

	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/auth"
)

// AuthUsecase is an interface for auth usecase
type AuthUsecase interface {
	Login(ctx context.Context, req presentations.LoginRequest) (*presentations.LoginResponse, error)
	Register(ctx context.Context, req presentations.RegisterRequest) (*presentations.RegisterResponse, error)
}

func NewAuthUsecase(
	userRepo repository.UserRepo,
	authenticator auth.Authenticator,
) AuthUsecase {
	return &authUsecase{
		userRepo:      userRepo,
		authenticator: authenticator,
	}
}

type authUsecase struct {
	userRepo      repository.UserRepo
	authenticator auth.Authenticator
}
