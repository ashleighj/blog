package presentation

import (
	"api/config"
	"api/logger"
	"context"
	"log"
	"net/http"
)

// API holds all
type API struct {
	conf     *config.Service
	router   *Router
	handlers []Handler // v1.Handler, v2.Handler etc
}

// StartServer starts the web server so that it is ready to handle
// incoming connections
func (a *API) StartServer(ctx context.Context) {
	portToListen := ":" + a.conf.Port

	logger.Infof(ctx, "server running on port: %d", a.conf.Port)
	log.Fatal(
		http.ListenAndServe(portToListen, a.router.r),
	)
}

// Handler is an interface that
// each api version's handler should implement
type Handler interface {
	InitRoutes(*Router)
}

// New constructs and returns a new API struct
func New(
	ctx context.Context,
	conf *config.Service,
	r *Router,
	handlers []Handler) *API {

	api := API{
		router:   r,
		conf:     conf,
		handlers: handlers}

	return &api
}
