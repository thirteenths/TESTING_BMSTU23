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

type UserHandler struct {
	r    handler.Renderer
	user app.UserService
}

func NewUserHandler(r handler.Renderer, user app.UserService) *UserHandler {
	return &UserHandler{
		r:    r,
		user: user,
	}
}

func (h *UserHandler) BasePrefix() string { return "/users" }

func (h *UserHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/{id}", h.r.Wrap(h.GetUser))
	r.Post("/login", h.r.Wrap(h.LogIn))
	// r.Get("/generate", h.r.Wrap(h.GenerateCode))
	r.Post("/verify", h.r.Wrap(h.Verify))
	r.Post("/update", h.r.Wrap(h.UpdatePassword))

	return r
}

func (h *UserHandler) GetUser(w http.ResponseWriter, req *http.Request) handler.Response {
	id, err := requests.NewGetEvent().Bind(req)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetUser GetFeed")
		return handler.BadRequestResponse()
	}

	res, err := h.user.GetUser(context.Background(), id)
	if err != nil {
		log.WithError(err).Warnf("can't service.GetUser GetEvent")
		return handler.InternalServerErrorResponse()
	}

	return handler.OkResponse(res)
}

func (h *UserHandler) Registry(w http.ResponseWriter, req *http.Request) handler.Response {
	return handler.OkResponse("")
}

func (h *UserHandler) LogIn(w http.ResponseWriter, req *http.Request) handler.Response {
	login := &requests.LoginUser{}
	err := json.NewDecoder(req.Body).Decode(login)
	if err != nil {
		log.WithError(err).Warnf("bad request LoginUser")
		return handler.BadRequestResponse()
	}

	err = h.user.CheckPassword(context.Background(), *login)
	if err != nil {
		log.WithError(err).Warnf(" LoginUser")
		return handler.BadRequestResponse()
	}

	err = h.user.SendCode(context.Background(), login.Email)
	if err != nil {
		log.WithError(err).Warnf(" LoginUser")
		return handler.BadRequestResponse()
	}

	return handler.OkResponse("Code send on email")
}

func (h *UserHandler) UpdatePassword(w http.ResponseWriter, req *http.Request) handler.Response {
	pass := &requests.UpdatePassword{}
	err := json.NewDecoder(req.Body).Decode(pass)
	if err != nil {
		log.WithError(err).Warnf("bad request LoginUser")
		return handler.BadRequestResponse()
	}

	err = h.user.UpdatePassword(context.Background(), *pass)
	if err != nil {
		log.WithError(err).Warnf(" LoginUser")
		return handler.BadRequestResponse()
	}

	err = h.user.SendCode(context.Background(), pass.Email)
	if err != nil {
		log.WithError(err).Warnf(" LoginUser")
		return handler.BadRequestResponse()
	}

	return handler.OkResponse("Code send on email")
}

func (h *UserHandler) Verify(w http.ResponseWriter, req *http.Request) handler.Response {
	code := &requests.VerifyCode{}
	err := json.NewDecoder(req.Body).Decode(code)
	if err != nil {
		log.WithError(err).Warnf("bad request VerifyCode")
		return handler.BadRequestResponse()
	}

	err = h.user.CheckCode(context.Background(), code.Code)
	if err != nil {
		log.WithError(err).Warnf("CodeUser")
		return handler.BadRequestResponse()
	}

	return handler.OkResponse("Code successful")
}
