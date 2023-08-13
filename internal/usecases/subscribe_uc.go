package usecases

import (
	"context"

	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/presentations"
)

// SubscribeUsecase is an interface for auth usecase
type SubscribeUsecase interface {
	Subscribe(ctx context.Context) (*presentations.SubscribeResponse, error)
}

func NewSubscribeUsecase(
	userRepo repository.UserRepo,
	userSubscriptionRepo repository.UserSubscriptionRepo,
) SubscribeUsecase {
	return &subscribeUsecase{
		userRepo:             userRepo,
		userSubscriptionRepo: userSubscriptionRepo,
	}
}

type subscribeUsecase struct {
	userRepo             repository.UserRepo
	userSubscriptionRepo repository.UserSubscriptionRepo
}
