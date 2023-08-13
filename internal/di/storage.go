package di

import "github.com/yonisaka/dating-service/pkg/storage"

func GetCloudinaryStorage() storage.Storage {
	cfg := GetConfig()
	return storage.NewCloudinaryClient(cfg.Cloudinary)
}
