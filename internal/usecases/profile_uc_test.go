package usecases_test

import (
	"github.com/yonisaka/dating-service/config"
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/usecases"
	"github.com/yonisaka/dating-service/pkg/storage"
)

type profileFields struct {
	cfg                  *config.Config
	userRepo             repository.UserRepo
	userSubscriptionRepo repository.UserSubscriptionRepo
	userImageRepo        repository.UserImageRepo
	userPreferenceRepo   repository.UserPreferenceRepo
	storage              storage.Storage
}

func profileSut(f profileFields) usecases.ProfileUsecase {
	return usecases.NewProfileUsecase(
		f.cfg,
		f.userRepo,
		f.userSubscriptionRepo,
		f.userImageRepo,
		f.userPreferenceRepo,
		f.storage,
	)
}
