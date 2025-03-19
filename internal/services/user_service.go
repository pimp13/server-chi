package services

import (
	"context"

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

func (service *UserService) Create(ctx context.Context, user *models.User) error {
	hash, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return service.repo.CreateNewUser(ctx, user)
}
