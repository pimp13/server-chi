package services

import (
	"context"
	"errors"

	"github.com/pimp13/server-chi/internal/models"
	"github.com/pimp13/server-chi/internal/repositories"
	"github.com/pimp13/server-chi/pkg/util"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, user *models.User) error {
	// check user exists and registered
	exists, err := s.repo.UserExistsByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user with this email already exists")
	}

	// validation data

	// hash password
	hash, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	return s.repo.Create(ctx, user)
}
