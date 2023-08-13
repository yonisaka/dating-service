package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/yonisaka/dating-service/internal/consts"
	"github.com/yonisaka/dating-service/pkg/auth"
)

type Authenticator interface {
	Authenticate(r *http.Request) int
}

type authenticator struct {
	jwt auth.Authenticator
}

func NewAuthenticator(jwt auth.Authenticator) Authenticator {
	return &authenticator{
		jwt: jwt,
	}
}

func (a *authenticator) Authenticate(r *http.Request) int {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return http.StatusUnauthorized
	}

	if !strings.Contains(authHeader, "Bearer") {
		return http.StatusUnauthorized
	}

	_, claims, err := a.jwt.RequestTokenJwt(authHeader)

	if err != nil {
		return http.StatusUnauthorized
	}

	req := r.WithContext(context.WithValue(r.Context(), consts.CtxAuthInfo, claims))
	*r = *req

	return http.StatusOK
}
