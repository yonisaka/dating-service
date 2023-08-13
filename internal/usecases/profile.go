package usecases

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/yonisaka/dating-service/internal/consts"
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/helper"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/logger"
	"github.com/yonisaka/dating-service/pkg/util"
	"golang.org/x/sync/errgroup"
)

// GetProfile is a function to get profile
//
//nolint:funlen
func (u *profileUsecase) GetProfile(ctx context.Context) (*presentations.ProfileResponse, error) {
	var (
		lf = logger.NewFields(
			logger.EventName("usecase.profile"),
		)
		authInfo     = helper.AuthInfoFromContext(ctx)
		user         *repository.User
		subscription *repository.UserSubscription
		userImages   []repository.UserImage
	)

	group, _ := errgroup.WithContext(ctx)

	group.Go(func() error {
		result, err := u.userRepo.FindByID(ctx, authInfo.UserID)
		if err != nil {
			return err
		}

		user = result
		return nil
	})

	group.Go(func() error {
		result, err := u.userSubscriptionRepo.FindByUserID(ctx, authInfo.UserID)
		if err != nil {
			return err
		}

		subscription = result
		return nil
	})

	group.Go(func() error {
		result, err := u.userImageRepo.FindByUserID(ctx, authInfo.UserID)
		if err != nil {
			return err
		}

		userImages = result
		return nil
	})

	if err := group.Wait(); err != nil {
		return nil, err
	}

	result := presentations.ProfileResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    consts.Member,
		Email:     user.Email,
		Phone:     user.Phone,
		Dob:       user.Dob.Format(time.DateTime),
		Gender:    user.Gender,
		Intend:    user.Intend,
	}

	if subscription != nil {
		result.Status = consts.Verified
	}

	if len(userImages) > 0 {
		for _, image := range userImages {
			signedURL, err := u.storage.GetSignedURL(u.cfg.Cloudinary.Bucket, image.ImageURL)
			if err != nil {
				logger.ErrorWithContext(ctx, "error get signed url", lf...)
				return nil, err
			}

			result.Images = append(result.Images, signedURL.String())
		}
	}

	return &result, nil
}

// UploadImages is a function to upload images
func (u *profileUsecase) UploadImages(ctx context.Context, req presentations.UploadImageRequest) error {
	var (
		lf = logger.NewFields(
			logger.EventName("usecase.upload_images"),
		)
		authInfo   = helper.AuthInfoFromContext(ctx)
		userImages = make([]repository.UserImage, 0)
	)

	for _, image := range req.Images {
		i := strings.Index(image, ",")
		if i < 0 {
			logger.ErrorWithContext(ctx, "invalid image", lf...)
			return fmt.Errorf("invalid image")
		}

		extIndex := strings.Index(image[:i], ";")
		if extIndex < 0 {
			logger.ErrorWithContext(ctx, "invalid image", lf...)
			return fmt.Errorf("invalid image")
		}

		mimeType := strings.Split(image[:extIndex], ":")[1]
		if !util.InArray(mimeType, consts.ExtImageBase64Allowed) {
			logger.ErrorWithContext(ctx, "invalid image extension", lf...)
			return fmt.Errorf("invalid image extension")
		}

		byteImage, err := base64.StdEncoding.DecodeString(image[i+1:])

		if err != nil {
			logger.ErrorWithContext(ctx, err.Error(), lf...)
			return err
		}

		filename := uuid.New().String()
		err = u.storage.Put(ctx, u.cfg.Cloudinary.Bucket, filename, byteImage, false, mimeType)

		if err != nil {
			logger.ErrorWithContext(ctx, err.Error(), lf...)
			return err
		}

		userImages = append(userImages, repository.UserImage{
			UserID:   authInfo.UserID,
			ImageURL: filename,
		})
	}

	err := u.userImageRepo.StoreBulk(ctx, userImages)
	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return err
	}

	return nil
}

// SetPreference is a function to set preference
func (u *profileUsecase) SetPreference(ctx context.Context, req presentations.UserPreferenceRequest) error {
	var (
		lf = logger.NewFields(
			logger.EventName("usecase.upload_images"),
		)
		authInfo = helper.AuthInfoFromContext(ctx)
	)

	err := u.userPreferenceRepo.Update(ctx, repository.UserPreference{
		UserID:    authInfo.UserID,
		MinAge:    req.MinAge,
		MaxAge:    req.MaxAge,
		UseIntend: req.UseIntend,
	})

	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return err
	}

	return nil
}
