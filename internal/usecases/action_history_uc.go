package usecases

import (
	"context"

	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/presentations"
)

// ActionHistoryUsecase is an interface for auth usecase
type ActionHistoryUsecase interface {
	GetActionHistories(ctx context.Context) ([]presentations.ActionHistoryResponse, error)
}

func NewActionHistoryUsecase(
	userActionHistoryRepo repository.UserActionHistoryRepo,
	userRepo repository.UserRepo,
) ActionHistoryUsecase {
	return &actionHistoryUsecase{
		userActionHistoryRepo: userActionHistoryRepo,
		userRepo:              userRepo,
	}
}

type actionHistoryUsecase struct {
	userActionHistoryRepo repository.UserActionHistoryRepo
	userRepo              repository.UserRepo
}
