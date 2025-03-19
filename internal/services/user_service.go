package services

import (
	"context"
	"github.com/pimp13/server-chi/internal/models"
	"github.com/pimp13/server-chi/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (service *UserService) Create(ctx context.Context, user *models.User) error {
	return service.repo.CreateNewUser(ctx, user)
}
