package types

import (
	"context"
	"github.com/pimp13/server-chi/internal/models"
)

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	CreateNewUser(ctx context.Context, user *models.User) error
}

type RegisterUserData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
