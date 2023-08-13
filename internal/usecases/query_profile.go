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
//nolint:all
func (u *queryProfileUsecase) GetQueryProfile(
	ctx context.Context,
	req presentations.QueryProfileRequest,
) (*presentations.QueryProfileResponse, error) {
	var (
		lf = logger.NewFields(
			logger.EventName("usecase.query.profile"),
		)
		intend           = make(chan string)
		userPreference   repository.UserPreference
		userSubscription *repository.UserSubscription
		authInfo         = helper.AuthInfoFromContext(ctx)
		keyUserQuery     = fmt.Sprintf("user:query:%d", authInfo.UserID)
		keyRecentQuery   = fmt.Sprintf("user:recent:query:%d", authInfo.UserID)
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

	recentQuery, err := u.kvs.Get(ctx, keyRecentQuery)
	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	if recentQuery != nil {
		var recentQueryUserID int64
		err = json.Unmarshal([]byte(recentQuery.(string)), &recentQueryUserID)

		if err != nil {
			logger.ErrorWithContext(ctx, err.Error(), lf...)
			return nil, err
		}

		if req.Action == "" {
			return nil, fmt.Errorf("you need to specify action")
		}

		err = u.userActionHistoryRepo.Store(ctx, repository.UserActionHistory{
			UserID:    authInfo.UserID,
			ProfileID: recentQueryUserID,
			Action:    req.Action,
		})

		if err != nil {
			logger.ErrorWithContext(ctx, err.Error(), lf...)
			return nil, err
		}
	}

	existQuery, err := u.kvs.Get(ctx, keyUserQuery)

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

		byteExistUserIDs, err := json.Marshal(existUserIDs)

		if err != nil {
			logger.ErrorWithContext(ctx, err.Error(), lf...)
			return nil, err
		}

		_, err = u.kvs.Set(ctx, keyUserQuery, byteExistUserIDs, time.Duration(24)*time.Hour)

		if err != nil {
			logger.ErrorWithContext(ctx, err.Error(), lf...)
			return nil, err
		}

		byteUserID, err := json.Marshal(user.ID)

		if err != nil {
			logger.ErrorWithContext(ctx, err.Error(), lf...)
			return nil, err
		}

		_, err = u.kvs.Set(ctx, keyRecentQuery, byteUserID, time.Duration(24)*time.Hour)

		if err != nil {
			logger.ErrorWithContext(ctx, err.Error(), lf...)
			return nil, err
		}

		profileImages, err := u.userImageRepo.FindByUserID(ctx, user.ID)
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

		if len(profileImages) > 0 {
			for _, image := range profileImages {
				signedURL, err := u.storage.GetSignedURL(u.cfg.Cloudinary.Bucket, image.ImageURL)
				if err != nil {
					logger.ErrorWithContext(ctx, "error get signed url", lf...)
					return nil, err
				}

				result.Images = append(result.Images, signedURL.String())
			}
		}

		break
	}

	if result == nil {
		return nil, fmt.Errorf("no more user match with your preference")
	}

	return result, nil
}
