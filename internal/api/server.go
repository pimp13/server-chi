package api

import (
	"database/sql"
	"github.com/gorilla/csrf"
	"log"
	"net/http"
	"time"

	"github.com/pimp13/server-chi/internal/handlers/user"
	"github.com/pimp13/server-chi/internal/repositories"
	"github.com/pimp13/server-chi/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/pimp13/server-chi/internal/handlers"
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
	s.middlewares(router)

	// Routes
	s.initRoutes(router)

	log.Printf("Server is running on: http://localhost%s", s.addr)
	return http.ListenAndServe(s.addr, router)
}

func (s *Server) initRoutes(router *chi.Mux) {
	// Routes
	up := handlers.NewHandler()

	// Register Services
	userRepo := repositories.NewUserRepository(s.db)
	userService := services.NewUserService(userRepo)
	userHandlers := user.NewUserHandler(userService)

	router.Route("/api", func(r chi.Router) {
		up.Routes(r)
		userHandlers.Routes(r)
	})
}

func (s *Server) middlewares(r *chi.Mux) {
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN", "X-XSRF-TOKEN"},
		ExposedHeaders:   []string{"Link", "X-CSRF-TOKEN"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	// CSRF token middleware
	csrfKey := []byte("32-byte-long-secret-key-5587996")
	csrfMiddleware := csrf.Protect(
		csrfKey,
		csrf.Secure(false),
		csrf.HttpOnly(true),
		csrf.SameSite(csrf.SameSiteStrictMode),
		csrf.CookieName("csrf_token"),
		csrf.Path("/"),
		csrf.MaxAge(86400),
	)
	r.Use(csrfMiddleware)
}
