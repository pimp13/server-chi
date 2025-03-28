package interfaces

import (
	"context"
	"github.com/pimp13/server-chi/internal/models"
)

type IUserRepository interface {
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id int) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	GetLatestAll(ctx context.Context) ([]models.User, error)
}
