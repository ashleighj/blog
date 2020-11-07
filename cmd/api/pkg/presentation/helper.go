package presentation

import (
	lg "api/logger"
	err "api/service_error"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// httpResponse standard json response
type successResponse struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Errors []err.ServiceError `json:"errors"`
}

// DecodeBody decodes a json request body into a given struct
func DecodeBody(ctx context.Context,
	data io.ReadCloser, model interface{}) (e error) {

	e = json.NewDecoder(data).Decode(model)
	if e != nil {
		lg.Error(ctx, e)
		return err.InputMalformed
	}

	return
}

// Respond sends standard json response
func Respond(ctx context.Context,
	w http.ResponseWriter, data interface{}, httpStatusCode int, e error) {

	if e != nil {
		RespondError(ctx, w, httpStatusCode, e)
		return
	}

	RespondSuccess(ctx, w, data)
}

// RespondSuccess sends a json success response
func RespondSuccess(
	ctx context.Context, w http.ResponseWriter, data interface{}) {

	id, ok := ctx.Value("x-request-id").(string)
	if !ok {
		id = ""
	}

	response := &successResponse{Data: data}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-Id", id)
	json.NewEncoder(w).Encode(response)
}

// RespondError sends a json error response
func RespondError(ctx context.Context,
	w http.ResponseWriter, httpStatusCode int, e error) {

	id, ok := ctx.Value("x-request-id").(string)
	if !ok {
		id = ""
	}

	var response errorResponse
	if er, ok := e.(err.ServiceError); ok {
		response.Errors = append(response.Errors, er)
	} else {
		response.Errors = append(response.Errors, err.Internal(e))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-Id", id)
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(response)
}
