package di

import (
	"github.com/yonisaka/dating-service/pkg/auth"
)

func GetAuthenticator() auth.Authenticator {
	return auth.NewAuthenticator(GetConfig())
}
