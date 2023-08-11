package di

import "github.com/yonisaka/dating-service/internal/usecases"

func GetHealthUsecase() usecases.HealthUsecase {
	return usecases.NewHealthUsecase(GetHealthRepo())
}
