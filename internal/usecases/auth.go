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

// Login is a function to login user
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

	ok, err := crypto.CompareHash(user.Password, req.Password)
	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	if !ok {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	userToken := u.authenticator.CreateTokenJWT(user.ID)

	result = presentations.LoginResponse{
		AccessToken:  userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
		ExpiredAt:    userToken.AtExpires.Format(time.DateTime),
	}

	return &result, nil
}

// Register is a function to register user
//
//nolint:funlen
func (u *authUsecase) Register(ctx context.Context, req presentations.RegisterRequest) (*presentations.RegisterResponse, error) { //nolint:lll
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
			Dob:       user.Dob.Format(time.DateTime),
			Gender:    user.Gender,
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
