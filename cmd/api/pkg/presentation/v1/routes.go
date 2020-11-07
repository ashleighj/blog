package v1

import (
	"api/pkg/presentation"
	"net/http"
)

// InitRoutes registers v1 routes and their handlers
func (h *V1Handler) InitRoutes(r *presentation.Router) {
	r.RouteGET("/v1/content/types", h.getContentTypes)
	r.RoutePUT("/v1/content/types", h.addContentType)
	r.RouteGET("/v1/content/articles", h.getArticles)
	r.RouteGET("/v1/contributors", h.getContributors)
}

func (h *V1Handler) getArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("articles"))
}

func (h *V1Handler) getContributors(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("contributors"))
}

func (h *V1Handler) addContentType(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("content type added"))
}

func (h *V1Handler) getContentTypes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("content types"))
}
