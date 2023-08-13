package di

import "github.com/yonisaka/dating-service/internal/middleware"

func GetAuthMiddleware() middleware.Authenticator {
	return middleware.NewAuthenticator(GetAuthenticator())
}
