package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Routes(r chi.Router) {
	r.Get("/user", h.getUser)
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world this is user handler"))
}
