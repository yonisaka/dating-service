package usecases

import (
	"context"

	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/kvs"
)

// QueryProfileUsecase is an interface for auth usecase
type QueryProfileUsecase interface {
	GetQueryProfile(ctx context.Context) (*presentations.QueryProfileResponse, error)
}

func NewQueryProfileUsecase(
	userRepo repository.UserRepo,
	userPreferenceRepo repository.UserPreferenceRepo,
	userSubscriptionRepo repository.UserSubscriptionRepo,
	kvs kvs.Client,
) QueryProfileUsecase {
	return &queryProfileUsecase{
		userRepo:             userRepo,
		userPreferenceRepo:   userPreferenceRepo,
		userSubscriptionRepo: userSubscriptionRepo,
		kvs:                  kvs,
	}
}

type queryProfileUsecase struct {
	userRepo             repository.UserRepo
	userPreferenceRepo   repository.UserPreferenceRepo
	userSubscriptionRepo repository.UserSubscriptionRepo
	kvs                  kvs.Client
}
