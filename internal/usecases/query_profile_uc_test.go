package usecases_test

import (
	"github.com/yonisaka/dating-service/config"
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/usecases"
	"github.com/yonisaka/dating-service/pkg/kvs"
	"github.com/yonisaka/dating-service/pkg/storage"
)

type queryProfileFields struct {
	cfg                   *config.Config
	userRepo              repository.UserRepo
	userPreferenceRepo    repository.UserPreferenceRepo
	userSubscriptionRepo  repository.UserSubscriptionRepo
	userActionHistoryRepo repository.UserActionHistoryRepo
	userImageRepo         repository.UserImageRepo
	kvs                   kvs.Client
	storage               storage.Storage
}

func queryProfileSut(f queryProfileFields) usecases.QueryProfileUsecase {
	return usecases.NewQueryProfileUsecase(
		f.cfg,
		f.userRepo,
		f.userPreferenceRepo,
		f.userSubscriptionRepo,
		f.userActionHistoryRepo,
		f.userImageRepo,
		f.kvs,
		f.storage,
	)
}
