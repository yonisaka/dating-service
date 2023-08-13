package di

import "github.com/yonisaka/dating-service/internal/adapters/httphandler"

// GetHealthHandler is a function to get health handler
func GetHealthHandler() httphandler.Handler {
	return httphandler.NewHealthHandler(GetHealthUsecase())
}

// GetRegisterHandler is a function to get register handler
func GetRegisterHandler() httphandler.Handler {
	return httphandler.NewRegisterHandler(GetAuthUsecase())
}

// GetLoginHandler is a function to get login handler
func GetLoginHandler() httphandler.Handler {
	return httphandler.NewLoginHandler(GetAuthUsecase())
}

// GetProfileHandler is a function to get profile handler
func GetProfileHandler() httphandler.Handler {
	return httphandler.NewProfileHandler(GetProfileUsecase())
}

// GetQueryProfileHandler is a function to get query profile handler
func GetQueryProfileHandler() httphandler.Handler {
	return httphandler.NewQueryProfileHandler(GetQueryProfileUsecase())
}

// GetSubscribeHandler is a function to get subscribe handler
func GetSubscribeHandler() httphandler.Handler {
	return httphandler.NewSubscribeHandler(GetSubscribeUsecase())
}

// GetUploadImageHandler is a function to get upload image handler
func GetUploadImageHandler() httphandler.Handler {
	return httphandler.NewUploadImageHandler(GetProfileUsecase())
}

// GetPreferenceHandler is a function to get preference handler
func GetPreferenceHandler() httphandler.Handler {
	return httphandler.NewPreferenceHandler(GetProfileUsecase())
}

// GetActionHistoryHandler is a function to get action history handler
func GetActionHistoryHandler() httphandler.Handler {
	return httphandler.NewActionHistoryHandler(GetActionHistoryUsecase())
}
