package user

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/pimp13/server-chi/internal/models"
	"github.com/pimp13/server-chi/internal/services"
	"github.com/pimp13/server-chi/pkg/requests"
	"github.com/pimp13/server-chi/pkg/util"
	"log"
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
	r.Get("/users", h.getAllUser)
	r.Post("/register", h.register)
	r.Post("/login", h.login)

	r.Post("/test-cors", h.testCors)
}

/* Handlers */
func (h *UserHandler) getAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.LatestAllUser(context.Background())
	if err != nil {
		_ = util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	_ = util.WriteJSON(w, http.StatusOK, users)
}

func (h *UserHandler) testCors(w http.ResponseWriter, r *http.Request) {
	_ = util.WriteJSON(w, http.StatusOK, map[string]string{"message": "Sending POST request form next for go cors"})
}

func (h *UserHandler) register(w http.ResponseWriter, r *http.Request) {
	// get user data and parse to json
	var request requests.UserRegisterRequest
	if err := util.ParseJSON(r, &request); err != nil {
		_ = util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate data
	if err := util.Validate.Struct(&request); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			_ = util.WriteError(w, http.StatusBadRequest, validationErrors)
		}
		return
	}

	// create user
	err := h.userService.RegisterUser(context.Background(), &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		_ = util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// send response
	if err := util.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": "user registered and created successfully",
	}); err != nil {
		log.Fatal(err)
		return
	}
}

func (h *UserHandler) login(w http.ResponseWriter, r *http.Request) {
	var request requests.UserLoginRequest
	if err := util.ParseJSON(r, &request); err != nil {
		_ = util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := util.Validate.Struct(&request); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			_ = util.WriteError(w, http.StatusBadRequest, validationErrors)
		}
		return
	}

	token, err := h.userService.LoginUser(context.Background(), request)
	if err != nil {
		_ = util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    token,
		Path:     "/",
		MaxAge:   86400,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	_ = util.WriteJSON(w, http.StatusOK, map[string]string{
		"token":   token,
		"message": "user is logged",
	})
}
