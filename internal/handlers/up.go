package handlers

import (
	"github.com/gorilla/csrf"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type IHandler interface {
	Routes(r chi.Router)
}

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Routes(router chi.Router) {
	router.Get("/up", h.upApplication)
	router.Get("/csrf-token", h.csrfToken)
}

func (h *Handler) upApplication(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application is up..."))
}

func (h *Handler) csrfToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-CSRF-TOKEN", csrf.Token(r))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("CSRF token is set in headers"))
}
