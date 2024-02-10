package http

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/thirteenths/test-bmstu23/internal/app"
	"github.com/thirteenths/test-bmstu23/pkg/handler"
)

type APIHandler struct {
	r   handler.Renderer
	api app.API
}

func NewAPIHandler(r handler.Renderer, api app.API) *APIHandler {
	return &APIHandler{
		r:   r,
		api: api,
	}
}

func (h *APIHandler) BasePrefix() string {
	return "/api"
}

func (h *APIHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/echo", h.r.Wrap(h.echo))

	return r
}

func (h *APIHandler) echo(_ http.ResponseWriter, r *http.Request) handler.Response {
	out := h.api.Echo()

	return handler.OkResponse(out)
}
