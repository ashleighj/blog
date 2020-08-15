package v1

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	router *mux.Router
	// service *service.Service
}

func newHandler() (*Handler, error) {
	h := Handler{}
	h.RegisterRoutes()
	return &h, nil
}

func (h *Handler) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("healthy"))
}
