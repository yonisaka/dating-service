package usecases_test

import (
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/usecases"
	"github.com/yonisaka/dating-service/pkg/auth"
)

type authFields struct {
	userRepo      repository.UserRepo
	authenticator auth.Authenticator
}

func authSut(f authFields) usecases.AuthUsecase {
	return usecases.NewAuthUsecase(
		f.userRepo,
		f.authenticator,
	)
}
