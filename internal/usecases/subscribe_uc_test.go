package usecases_test

import (
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/usecases"
)

type subscribeFields struct {
	userRepo             repository.UserRepo
	userSubscriptionRepo repository.UserSubscriptionRepo
}

func subscribeSut(f subscribeFields) usecases.SubscribeUsecase {
	return usecases.NewSubscribeUsecase(
		f.userRepo,
		f.userSubscriptionRepo,
	)
}
