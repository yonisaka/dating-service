package di

import (
	"net/http"

	"github.com/yonisaka/dating-service/config"
	"github.com/yonisaka/dating-service/pkg/routerkit"
)

type router struct {
	cfg    *config.Config
	router *routerkit.Router
}

func NewRouter() Router {
	cfg := GetConfig()

	RegistryMessage() //nolint:nolint

	return &router{cfg: cfg, router: routerkit.NewRouter(routerkit.WithServiceName(cfg.App.Name))}
}

func (r *router) Route() *routerkit.Router {
	root := r.router.PathPrefix("/").Subrouter()

	root.HandleFunc("/liveness", r.handle(
		httpGateway,
		GetHealthHandler(),
	)).Methods(http.MethodGet)

	api := root.PathPrefix("/api").Subrouter()
	apiV1 := api.PathPrefix("/v1").Subrouter()

	apiV1.HandleFunc("/register", r.handle(
		httpGateway,
		GetRegisterHandler(),
	)).Methods(http.MethodPost)

	apiV1.HandleFunc("/login", r.handle(
		httpGateway,
		GetLoginHandler(),
	)).Methods(http.MethodPost)

	apiV1.HandleFunc("/profile", r.handle(
		httpGateway,
		GetProfileHandler(),
		GetAuthMiddleware().Authenticate,
	)).Methods(http.MethodGet)

	apiV1.HandleFunc("/query-profile", r.handle(
		httpGateway,
		GetQueryProfileHandler(),
		GetAuthMiddleware().Authenticate,
	)).Methods(http.MethodGet)

	apiV1.HandleFunc("/subscribe", r.handle(
		httpGateway,
		GetSubscribeHandler(),
		GetAuthMiddleware().Authenticate,
	)).Methods(http.MethodPost)

	apiV1.HandleFunc("/upload-image", r.handle(
		httpGateway,
		GetUploadImageHandler(),
		GetAuthMiddleware().Authenticate,
	)).Methods(http.MethodPost)

	apiV1.HandleFunc("/preference", r.handle(
		httpGateway,
		GetPreferenceHandler(),
		GetAuthMiddleware().Authenticate,
	)).Methods(http.MethodPost)

	return r.router
}
