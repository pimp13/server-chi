package user

import (
	"context"
	"log"
	"net/http"

	"github.com/pimp13/server-chi/internal/models"
	"github.com/pimp13/server-chi/internal/services"
	"github.com/pimp13/server-chi/pkg/interfaces"
	"github.com/pimp13/server-chi/pkg/util"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Routes(r chi.Router) {
	r.Get("/user", h.getUser)
	r.Post("/register", h.register)

	r.Post("/test-cors", h.testCors)
}

/* Handlers */
func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	_ = util.WriteJSON(w, http.StatusOK, "hello world this is go server")
}

func (h *UserHandler) testCors(w http.ResponseWriter, r *http.Request) {
	_ = util.WriteJSON(w, http.StatusOK, map[string]string{"message": "Sending POST request form next for go cors"})
}

func (h *UserHandler) register(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// get user data and parse to json
	var request interfaces.UserRegisterRequest
	if err := util.ParseJSON(r, &request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// create user
	err := h.userService.RegisterUser(ctx, &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send response
	if err = util.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": "user registered and created successfully",
	}); err != nil {
		log.Fatal(err)
		return
	}
}
