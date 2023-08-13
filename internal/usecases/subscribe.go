package usecases

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/helper"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/logger"
)

func (u *subscribeUsecase) Subscribe(ctx context.Context) (*presentations.SubscribeResponse, error) {
	var (
		lf = logger.NewFields(
			logger.EventName("usecase.profile"),
		)
		authInfo = helper.AuthInfoFromContext(ctx)
	)

	user, err := u.userRepo.FindByID(ctx, authInfo.UserID)
	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	subscriptionCode := uuid.New().String()
	subscriptionExpiredAt := time.Now().AddDate(0, 1, 0)
	err = u.userSubscriptionRepo.Store(ctx, repository.UserSubscription{
		UserID:           user.ID,
		SubscriptionCode: subscriptionCode,
		ExpiredAt:        &subscriptionExpiredAt,
	})

	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	result := presentations.SubscribeResponse{
		SubscriptionCode: subscriptionCode,
		ExpiredAt:        subscriptionExpiredAt.Format(time.DateTime),
	}

	return &result, nil
}
