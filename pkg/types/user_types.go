package types

import (
	"context"
	"github.com/pimp13/server-chi/internal/models"
)

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateNewUser(user *models.User) error
}
