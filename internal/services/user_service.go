package services

import (
	"context"
	"errors"
	"github.com/pimp13/server-chi/internal/models"
	"github.com/pimp13/server-chi/internal/repositories"
	"github.com/pimp13/server-chi/pkg/auth"
	"github.com/pimp13/server-chi/pkg/requests"
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

	// hash password
	hash, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	return s.repo.Create(ctx, user)
}

func (s *UserService) LatestAllUser(ctx context.Context) ([]models.User, error) {
	users, err := s.repo.GetLatestAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) LoginUser(ctx context.Context, request requests.UserLoginRequest) (string, error) {
	user, err := s.repo.FindByEmail(ctx, request.Email)
	if err != nil {
		return "", err
	}

	if !util.CheckHashPassword(user.Password, request.Password) {
		return "", errors.New("password no match")
	}

	token, err := auth.MakeToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
