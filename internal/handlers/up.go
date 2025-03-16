package handlers

import (
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
}

func (h *Handler) upApplication(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application is up..."))
}
