package presentation

import (
	lg "api/logger"
	"context"
	"errors"
	"net/http"
	"os"
	"time"
)

var (
	// ErrBadPayload the error to be returned when a request is malformed
	ErrBadPayload = errors.New("Malformed json request payload")
)

// BaseHandler is the handler for base api routes
type BaseHandler struct {
}

// NewBaseHandler creates, intialises and returns a basic
// api route handler
func NewBaseHandler(ctx context.Context, r *Router) *BaseHandler {
	b := BaseHandler{}
	b.InitRoutes(r)
	return &b
}

// InitRoutes initializes base (non-versioned) api routes
func (h *BaseHandler) InitRoutes(r *Router) {
	r.RouteGET("/healthcheck", h.healthCheck)
}

// healthcheck response
type healthcheck struct {
	Host     string    `json:"host"`
	Datetime time.Time `json:"datetime"`
}

func (h *BaseHandler) healthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var out healthcheck
	var e error

	out.Host, e = os.Hostname()
	if e != nil {
		lg.Error(ctx, e)
		RespondError(ctx, w, http.StatusOK, e)
		return
	}

	out.Datetime = time.Now()
	Respond(ctx, w, out, http.StatusOK, e)
}
