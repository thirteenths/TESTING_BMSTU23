package http

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/thirteenths/test-bmstu23/internal/app"
	"github.com/thirteenths/test-bmstu23/internal/domain/requests"
	"github.com/thirteenths/test-bmstu23/pkg/handler"
)

type EventHandler struct {
	r     handler.Renderer
	event app.EventService
}

func NewEventHandler(r handler.Renderer, event app.EventService) *EventHandler {
	return &EventHandler{
		r:     r,
		event: event,
	}
}
func (h *EventHandler) BasePrefix() string { return "/events" }

func (h *EventHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.r.Wrap(h.GetAllEvent))
	r.Get("/{id}", h.r.Wrap(h.GetEvent))
	r.Delete("/{id}", h.r.Wrap(h.DeleteEvent))
	r.Post("/", h.r.Wrap(h.CreateEvent))

	return r
}

func (h *EventHandler) GetAllEvent(w http.ResponseWriter, _ *http.Request) handler.Response {
	res, err := h.event.GetAllEvent(context.Background())
	if err != nil {
		log.WithError(err).Warnf("can't service.GetAllEvent GetAllEvent")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}

func (h *EventHandler) GetEvent(w http.ResponseWriter, req *http.Request) handler.Response {
	id, err := requests.NewGetEvent().Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetFeed GetFeed")
		return handler.BadRequestResponse()
	}

	res, err := h.event.GetEvent(context.Background(), id)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetEvent GetEvent")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}

func (h *EventHandler) DeleteEvent(w http.ResponseWriter, req *http.Request) handler.Response {
	id, err := requests.NewGetEvent().Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetFeed GetFeed")
		return handler.BadRequestResponse()
	}

	err = h.event.DeleteEvent(context.Background(), id)

	if err != nil {
		log.WithError(err).Warnf("can't service.DeleteEvent DeleteEvent")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse("OK")
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, req *http.Request) handler.Response {
	event := &requests.CreateEvent{}
	err := json.NewDecoder(req.Body).Decode(event)
	if err != nil {
		log.WithError(err).Warnf("bad request CreateEvent")
		return handler.BadRequestResponse()
	}

	res, err := h.event.CreateEvent(context.Background(), *event)
	if err != nil {
		log.WithError(err).Warnf("can't service.CreateEvent CreateEvent")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}
