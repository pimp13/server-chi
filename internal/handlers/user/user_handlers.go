package user

import (
	"context"
	"github.com/pimp13/server-chi/internal/models"
	"github.com/pimp13/server-chi/internal/services"
	"github.com/pimp13/server-chi/pkg/types"
	"github.com/pimp13/server-chi/pkg/util"
	"net/http"

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
}

/* Handlers */
func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world this is user handler"))
}

func (h *UserHandler) register(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// get user data and parse to json
	var user types.RegisterUserData
	if err := util.ParseJSON(r, &user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// validation data

	// check user exists and registered

	// create user
	err := h.userService.Create(ctx, &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send response
	err = util.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user registered and created successfully"})
	if err != nil {
		return
	}
}
