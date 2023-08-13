package di

import "github.com/yonisaka/dating-service/internal/usecases"

func GetHealthUsecase() usecases.HealthUsecase {
	return usecases.NewHealthUsecase(GetHealthRepo())
}

func GetAuthUsecase() usecases.AuthUsecase {
	return usecases.NewAuthUsecase(GetUserRepo(), GetAuthenticator())
}

func GetProfileUsecase() usecases.ProfileUsecase {
	return usecases.NewProfileUsecase(
		GetConfig(), GetUserRepo(), GetUserSubscriptionRepo(), GetUserImageRepo(), GetCloudinaryStorage())
}

func GetQueryProfileUsecase() usecases.QueryProfileUsecase {
	return usecases.NewQueryProfileUsecase(GetUserRepo(), GetUserPreferenceRepo(), GetUserSubscriptionRepo(), GetRedis())
}

func GetSubscribeUsecase() usecases.SubscribeUsecase {
	return usecases.NewSubscribeUsecase(GetUserRepo(), GetUserSubscriptionRepo())
}
