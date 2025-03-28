package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/pimp13/server-chi/internal/models"
	"github.com/pimp13/server-chi/pkg/interfaces"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

var _ interfaces.IUserRepository = (*UserRepository)(nil)

func (repo *UserRepository) FindByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	query := `
			SELECT id, name, email, created_at
			FROM users
			WHERE id = ?
	`
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return &user, nil
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `
			SELECT id, name, email, password, created_at
			FROM users
			WHERE email = ?
	`
	err := repo.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return &user, nil
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

func (repo *UserRepository) GetLatestAll(ctx context.Context) ([]models.User, error) {
	query := `
			SELECT id, name, email, created_at
			FROM users
			ORDER BY created_at DESC
			LIMIT 10 OFFSET 0;
	`
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
