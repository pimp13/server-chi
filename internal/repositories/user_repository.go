package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/pimp13/server-chi/internal/models"
	"github.com/pimp13/server-chi/pkg/types"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

var _ types.UserRepositoryInterface = (*UserRepository)(nil)

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `
			SELECT id, name, email, password, created_at
			FROM users
			WHERE email = ?
	`
	err := repo.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return &user, nil
}
func (repo *UserRepository) GetUserByID(id uint) (*models.User, error) {
	panic("t")
}
func (repo *UserRepository) CreateNewUser(user *models.User) error {
	panic("t")
}
