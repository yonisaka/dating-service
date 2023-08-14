package usecases_test

import (
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/usecases"
)

type actionHistoryFields struct {
	userActionHistoryRepo repository.UserActionHistoryRepo
	userRepo              repository.UserRepo
}

func actionHistorySut(f actionHistoryFields) usecases.ActionHistoryUsecase {
	return usecases.NewActionHistoryUsecase(
		f.userActionHistoryRepo,
		f.userRepo,
	)
}
