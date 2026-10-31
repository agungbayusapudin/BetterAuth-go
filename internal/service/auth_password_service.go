package service

import (
	"better_auth/internal/adapter/repository/postgres"
	"better_auth/internal/core/domain"
)

type AuthPasswordService struct {
	repo postgres.UserRepository
}

func NewAuthPasswordService(userRepository postgres.UserRepository) *AuthPasswordService {
	return &AuthPasswordService{
		repo: userRepository,
	}
}

func (s *AuthPasswordService) Login(user *domain.User) (*domain.User, error) {
	return s.repo.GetUser(user.ID)
}

func (s *AuthPasswordService) Logout(id string) (bool, error) {
	return false, nil
}

func (s *AuthPasswordService) Register(user *domain.User) (*domain.User, error) {
	return s.repo.GetUser(user.ID)
}
