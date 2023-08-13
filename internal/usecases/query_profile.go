package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/helper"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/logger"
	"github.com/yonisaka/dating-service/pkg/util"
	"golang.org/x/sync/errgroup"
)

// GetQueryProfile QueryProfileUsecase is an interface for auth usecase
//
//nolint:funlen
func (u *queryProfileUsecase) GetQueryProfile(ctx context.Context) (*presentations.QueryProfileResponse, error) {
	var (
		lf = logger.NewFields(
			logger.EventName("usecase.query.profile"),
		)
		intend           = make(chan string)
		userPreference   repository.UserPreference
		userSubscription *repository.UserSubscription
		authInfo         = helper.AuthInfoFromContext(ctx)
		key              = fmt.Sprintf("user:query:%d", authInfo.UserID)
		result           *presentations.QueryProfileResponse
	)

	existUserIDs := make([]int64, 0)
	group, _ := errgroup.WithContext(ctx)

	group.Go(func() error {
		result, err := u.userRepo.FindByID(ctx, authInfo.UserID)
		if err != nil {
			return err
		}

		intend <- result.Intend

		userPreference.Gender = result.OppositeGender()

		return nil
	})

	group.Go(func() error {
		result, err := u.userPreferenceRepo.FindByUserID(ctx, authInfo.UserID)
		if err != nil {
			return err
		}

		if result != nil {
			userPreference = *result
			userPreference.Intend = <-intend
		}

		return nil
	})

	group.Go(func() error {
		result, err := u.userSubscriptionRepo.FindByUserID(ctx, authInfo.UserID)
		if err != nil {
			return err
		}

		userSubscription = result
		return nil
	})

	err := group.Wait()
	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	existQuery, err := u.kvs.Get(ctx, key)

	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	if existQuery != nil {
		err = json.Unmarshal([]byte(existQuery.(string)), &existUserIDs)
		if err != nil {
			logger.ErrorWithContext(ctx, err.Error(), lf...)
			return nil, err
		}

		if len(existUserIDs) > 10 && userSubscription == nil {
			return nil, fmt.Errorf("query profile is limited to 10 users, subscribe to premium to get more")
		}
	}

	users, err := u.userRepo.Find(ctx, userPreference)
	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	for _, user := range users {
		if len(existUserIDs) > 0 {
			if util.InArray(user.ID, existUserIDs) {
				continue
			}
		}

		existUserIDs = append(existUserIDs, user.ID)

		b, err := json.Marshal(existUserIDs)

		if err != nil {
			logger.ErrorWithContext(ctx, err.Error(), lf...)
			return nil, err
		}

		_, err = u.kvs.Set(ctx, key, b, time.Duration(24)*time.Hour)

		if err != nil {
			logger.ErrorWithContext(ctx, err.Error(), lf...)
			return nil, err
		}

		result = &presentations.QueryProfileResponse{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Age:       user.Age(),
			Gender:    user.Gender,
			Intend:    user.Intend,
		}

		break
	}

	if result == nil {
		return nil, fmt.Errorf("no user found")
	}

	return result, nil
}
