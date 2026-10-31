package postgres

import (
	"better_auth/internal/core/domain"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (r *UserRepository) UpdateUser(id string, user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (r *UserRepository) DeleteUser(id string) (bool, error) {
	return false, nil
}

func (r *UserRepository) GetUser(id string) (*domain.User, error) {
	return nil, nil
}

func (r *UserRepository) GetAllUser(id []string) (*domain.User, error) {
	return nil, nil
}
