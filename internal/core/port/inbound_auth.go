package port

import "better_auth/internal/core/domain"

type IAuthService interface {
	Login(user *domain.User) (*domain.User, error)
	Logout(id string) (*bool, error)
	Register(user *domain.User) (*domain.User, error)
}
