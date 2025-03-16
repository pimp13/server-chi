package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pimp13/server-chi/internal/handlers"
	"github.com/pimp13/server-chi/internal/handlers/user"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Start() error {
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Routes
	s.initRoutes(router)

	log.Printf("Server is running on: http://localhost%s", s.addr)
	return http.ListenAndServe(s.addr, router)
}

func (s *Server) initRoutes(router *chi.Mux) {
	// Routes
	handlers := handlers.NewHandler()

	userHandlers := user.NewHandler()
	router.Route("/api", func(r chi.Router) {
		handlers.Routes(r)
		userHandlers.Routes(r)
	})
}
