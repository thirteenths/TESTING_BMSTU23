package requests

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type GetEvent struct {
	ID int `json:"id"`
}

func NewGetEvent() *GetEvent {
	return &GetEvent{}
}

func (f *GetEvent) Bind(req *http.Request) (int, error) {
	return strconv.Atoi(chi.URLParam(req, "id"))
}
