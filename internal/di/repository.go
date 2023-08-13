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

// GetUserPreferenceRepo returns UserPreferenceRepo instance.
func GetUserPreferenceRepo() repository.UserPreferenceRepo {
	return datastore.NewUserPreferenceRepo(GetBaseRepo())
}

// GetUserSubscriptionRepo returns UserSubscriptionRepo instance.
func GetUserSubscriptionRepo() repository.UserSubscriptionRepo {
	return datastore.NewUserSubscriptionRepo(GetBaseRepo())
}

// GetUserImageRepo returns UserImageRepo instance.
func GetUserImageRepo() repository.UserImageRepo {
	return datastore.NewUserImageRepo(GetBaseRepo())
}
