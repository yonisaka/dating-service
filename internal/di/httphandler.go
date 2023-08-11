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
