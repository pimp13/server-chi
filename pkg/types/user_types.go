package types

import (
	"context"
	"github.com/pimp13/server-chi/internal/models"
)

type UserRepositoryInterface interface {
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id uint) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
