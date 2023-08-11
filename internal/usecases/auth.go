package usecases

import (
	"context"
	"time"

	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/crypto"
	"github.com/yonisaka/dating-service/pkg/logger"
	"golang.org/x/sync/errgroup"
)

func (u *authUsecase) Login(ctx context.Context, req presentations.LoginRequest) (*presentations.LoginResponse, error) {
	var (
		lf = logger.NewFields(
			logger.EventName("usecase.auth.login"),
		)
		result presentations.LoginResponse
	)

	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	userToken := u.token.CreateTokenJWT(user.ID)

	result = presentations.LoginResponse{
		AccessToken:  userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
		ExpiredAt:    userToken.AtExpires,
	}

	return &result, nil
}

func (u *authUsecase) Register(ctx context.Context, req presentations.RegisterRequest) (*presentations.RegisterResponse, error) {
	var (
		lf = logger.NewFields(
			logger.EventName("usecase.auth.register"),
		)
		result  presentations.RegisterResponse
		emailCh = make(chan string)
	)

	group, _ := errgroup.WithContext(ctx)

	group.Go(func() error {
		hashedPassword, err := crypto.Hash(req.Password)
		if err != nil {
			return err
		}

		dob, err := time.Parse("2006-01-02", req.Dob)
		if err != nil {
			return err
		}

		err = u.userRepo.Store(ctx, &repository.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
			Password:  hashedPassword,
			Dob:       &dob,
			Gender:    req.Gender,
			Intend:    req.Intend,
		})
		if err != nil {
			return err
		}

		emailCh <- req.Email

		return nil
	})

	group.Go(func() error {
		user, err := u.userRepo.FindByEmail(ctx, <-emailCh)
		if err != nil {
			return err
		}

		result = presentations.RegisterResponse{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
		}

		return nil
	})

	err := group.Wait()
	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	return &result, nil
}
