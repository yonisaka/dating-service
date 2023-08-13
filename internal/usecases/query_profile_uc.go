package usecases

import (
	"context"

	"github.com/yonisaka/dating-service/config"
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/kvs"
	"github.com/yonisaka/dating-service/pkg/storage"
)

// QueryProfileUsecase is an interface for auth usecase
type QueryProfileUsecase interface {
	GetQueryProfile(ctx context.Context, req presentations.QueryProfileRequest) (*presentations.QueryProfileResponse, error)
}

func NewQueryProfileUsecase(
	cfg *config.Config,
	userRepo repository.UserRepo,
	userPreferenceRepo repository.UserPreferenceRepo,
	userSubscriptionRepo repository.UserSubscriptionRepo,
	userActionHistoryRepo repository.UserActionHistoryRepo,
	userImageRepo repository.UserImageRepo,
	kvs kvs.Client,
	storage storage.Storage,
) QueryProfileUsecase {
	return &queryProfileUsecase{
		cfg:                   cfg,
		userRepo:              userRepo,
		userPreferenceRepo:    userPreferenceRepo,
		userSubscriptionRepo:  userSubscriptionRepo,
		userActionHistoryRepo: userActionHistoryRepo,
		userImageRepo:         userImageRepo,
		kvs:                   kvs,
		storage:               storage,
	}
}

type queryProfileUsecase struct {
	cfg                   *config.Config
	userRepo              repository.UserRepo
	userPreferenceRepo    repository.UserPreferenceRepo
	userSubscriptionRepo  repository.UserSubscriptionRepo
	userActionHistoryRepo repository.UserActionHistoryRepo
	userImageRepo         repository.UserImageRepo
	kvs                   kvs.Client
	storage               storage.Storage
}
