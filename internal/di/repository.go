package di

import (
	"github.com/yonisaka/dating-service/internal/entities/repository"
	"github.com/yonisaka/dating-service/internal/infrastructure/datastore"
)

// GetBaseRepo returns BaseRepo instance.
func GetBaseRepo() *datastore.BaseRepo {
	cfg := GetConfig()
	return datastore.NewBaseRepo(datastore.GetDatabaseMaster(&cfg.MasterDB), datastore.GetDatabaseSlave(&cfg.SlaveDB))
}

// GetHealthRepo returns HealthRepo instance.
func GetHealthRepo() repository.HealthRepo {
	return datastore.NewHealthRepo(GetBaseRepo())
}

// GetUserRepo returns UserRepo instance.
func GetUserRepo() repository.UserRepo {
	return datastore.NewUserRepo(GetBaseRepo())
}
