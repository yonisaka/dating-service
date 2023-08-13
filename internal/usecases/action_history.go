package usecases

import (
	"context"
	"time"

	"github.com/yonisaka/dating-service/internal/helper"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/logger"
)

func (u *actionHistoryUsecase) GetActionHistories(ctx context.Context) ([]presentations.ActionHistoryResponse, error) {
	var (
		lf = logger.NewFields(
			logger.EventName("usecase.action.history"),
		)
		authInfo = helper.AuthInfoFromContext(ctx)
		result   = make([]presentations.ActionHistoryResponse, 0)
	)

	actionHistories, err := u.userActionHistoryRepo.FindByUserID(ctx, authInfo.UserID)
	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	for _, actionHistory := range actionHistories {
		profile, err := u.userRepo.FindByID(ctx, actionHistory.ProfileID)
		if err != nil {
			logger.ErrorWithContext(ctx, err.Error(), lf...)
			return nil, err
		}

		result = append(result, presentations.ActionHistoryResponse{
			FirstName: profile.FirstName,
			LastName:  profile.LastName,
			Age:       profile.Age(),
			Action:    actionHistory.Action,
			UpdatedAt: actionHistory.UpdatedAt.Format(time.DateTime),
		})
	}

	return result, nil
}
