package presentation

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	r *httprouter.Router
}

// NewRouter creates and returns a new Router struct
func NewRouter() *Router {
	return &Router{
		r: httprouter.New()}
}

// RouteGET registers an http GET route
func (r *Router) RouteGET(route string, handler func(http.ResponseWriter, *http.Request)) {
	r.r.Handler(http.MethodGet, route, http.HandlerFunc(handler))
}

// RoutePOST registers an http POST route
func (r *Router) RoutePOST(route string, handler func(http.ResponseWriter, *http.Request)) {
	r.r.Handler(http.MethodPost, route, http.HandlerFunc(handler))
}

// RoutePUT registers an http PUT route
func (r *Router) RoutePUT(route string, handler func(http.ResponseWriter, *http.Request)) {
	r.r.Handler(http.MethodPut, route, http.HandlerFunc(handler))
}

// RouteDELETE registers an http DELETE route
func (r *Router) RouteDELETE(route string, handler func(http.ResponseWriter, *http.Request)) {
	r.r.Handler(http.MethodDelete, route, http.HandlerFunc(handler))
}

// RoutePATCH registers an http PATCH route
func (r *Router) RoutePATCH(route string, handler func(http.ResponseWriter, *http.Request)) {
	r.r.Handler(http.MethodPatch, route, http.HandlerFunc(handler))
}
