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

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
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
func (repo *UserRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	panic("find user by email")
}
func (repo *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (name, email, password) 
		VALUES (?, ?, ?)
 	`
	_, err := repo.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)
	return err
}

func (repo *UserRepository) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)`
	err := repo.db.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
