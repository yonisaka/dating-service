package usecases

import (
	"context"
	"time"

	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/presentations"
	"github.com/yonisaka/dating-service/pkg/crypto"
	"github.com/yonisaka/dating-service/pkg/logger"
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
func (u *authUsecase) Register(ctx context.Context, req presentations.RegisterRequest) (*presentations.RegisterResponse, error) {
	var (
		lf = logger.NewFields(
			logger.EventName("usecase.auth.register"),
		)
		result presentations.RegisterResponse
	)

	hashedPassword, err := crypto.Hash(req.Password)
	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
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
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		logger.ErrorWithContext(ctx, err.Error(), lf...)
		return nil, err
	}

	result = presentations.RegisterResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Dob:       user.Dob.Format(time.DateTime),
		Gender:    user.Gender,
	}

	return &result, nil
}
