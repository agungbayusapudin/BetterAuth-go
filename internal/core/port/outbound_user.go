package port

import "better_auth/internal/core/domain"

type OAuthRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	UpdateUser(id string, user *domain.User) (*domain.User, error)
	DeleteUser(id string) (*bool, error)
	GetUser(id string) (*domain.User, error)
	GetAllUser(id []string) (*domain.User, error)
}
