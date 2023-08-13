package usecases

import (
	"context"

	"github.com/yonisaka/dating-service/config"
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/storage"
)

// ProfileUsecase is an interface for auth usecase
type ProfileUsecase interface {
	GetProfile(ctx context.Context) (*presentations.ProfileResponse, error)
	UploadImages(ctx context.Context, req presentations.UploadImageRequest) error
	SetPreference(ctx context.Context, req presentations.UserPreferenceRequest) error
}

func NewProfileUsecase(
	cfg *config.Config,
	userRepo repository.UserRepo,
	userSubscriptionRepo repository.UserSubscriptionRepo,
	userImageRepo repository.UserImageRepo,
	userPreferenceRepo repository.UserPreferenceRepo,
	storage storage.Storage,
) ProfileUsecase {
	return &profileUsecase{
		cfg:                  cfg,
		userRepo:             userRepo,
		userSubscriptionRepo: userSubscriptionRepo,
		userImageRepo:        userImageRepo,
		userPreferenceRepo:   userPreferenceRepo,
		storage:              storage,
	}
}

type profileUsecase struct {
	cfg                  *config.Config
	userRepo             repository.UserRepo
	userSubscriptionRepo repository.UserSubscriptionRepo
	userImageRepo        repository.UserImageRepo
	userPreferenceRepo   repository.UserPreferenceRepo
	storage              storage.Storage
}
