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
	return &router{cfg: cfg, router: routerkit.NewRouter(routerkit.WithServiceName(cfg.App.Name))}
}

func (r *router) Route() *routerkit.Router {
	root := r.router.PathPrefix("/").Subrouter()

	healthHandler := GetHealthHandler()

	root.HandleFunc("/liveness", r.handle(
		httpGateway,
		healthHandler,
	)).Methods(http.MethodGet)

	return r.router
}
