package v1

import (
	"api/config"
	"api/pkg/blog/service"
	"api/pkg/presentation"
	"net/http"
)

type V1Handler struct {
	conf    *config.Config
	service *service.Service
}

// NewHandler returns a route handler for v1 of api functionality
func NewHandler(
	conf *config.Config,
	srv *service.Service,
	r *presentation.Router) (*V1Handler, error) {

	h := V1Handler{conf: conf, service: srv}
	h.InitRoutes(r)
	return &h, nil
}

func (h *V1Handler) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("healthy"))
}
