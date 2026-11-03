package postgres

import (
	"better_auth/internal/core/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// CreateUser membuat user baru dengan transaction context
func (r *UserRepository) CreateUser(user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	query := `
		INSERT INTO users (email, phone, display_name, avatar_url, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, email, phone, display_name, avatar_url, status, created_at, updated_at
	`

	now := time.Now()
	err = tx.QueryRowContext(ctx, query,
		user.Email,
		user.Phone,
		user.DisplayName,
		user.AvatarUrl,
		user.Status,
		now,
		now,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Phone,
		&user.DisplayName,
		&user.AvatarUrl,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("failed to create user: no rows returned")
		}
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return user, nil
}

// UpdateUser mengupdate user dengan transaction context
func (r *UserRepository) UpdateUser(id string, user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	query := `
		UPDATE users
		SET email = $1, phone = $2, display_name = $3, avatar_url = $4, status = $5, updated_at = $6
		WHERE id = $7 AND deleted_at IS NULL
		RETURNING id, email, phone, display_name, avatar_url, status, created_at, updated_at
	`

	now := time.Now()
	err = tx.QueryRowContext(ctx, query,
		user.Email,
		user.Phone,
		user.DisplayName,
		user.AvatarUrl,
		user.Status,
		now,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Phone,
		&user.DisplayName,
		&user.AvatarUrl,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found or already deleted")
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return user, nil
}

// DeleteUser melakukan soft delete pada user dengan transaction context
func (r *UserRepository) DeleteUser(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return false, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	query := `
		UPDATE users
		SET deleted_at = $1, updated_at = $2, status = $3
		WHERE id = $4 AND deleted_at IS NULL
	`

	now := time.Now()
	result, err := tx.ExecContext(ctx, query, now, now, "deleted", id)
	if err != nil {
		return false, fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return false, errors.New("user not found or already deleted")
	}

	if err = tx.Commit(); err != nil {
		return false, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return true, nil
}

// GetUser mengambil single user berdasarkan ID
func (r *UserRepository) GetUser(id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, email, email_verified_at, phone, phone_verified_at, display_name, avatar_url, status, created_at, updated_at
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`

	user := &domain.User{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.EmailVerifiedAt,
		&user.Phone,
		&user.PhoneVerifiedAt,
		&user.DisplayName,
		&user.AvatarUrl,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// GetAllUser mengambil multiple users berdasarkan list ID
func (r *UserRepository) GetAllUser(ids []string) (*domain.User, error) {
	if len(ids) == 0 {
		return nil, errors.New("ids list cannot be empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT id, email, email_verified_at, phone, phone_verified_at, display_name, avatar_url, status, created_at, updated_at
		FROM users
		WHERE id = ANY($1) AND deleted_at IS NULL
		ORDER BY created_at DESC
	`

	rows, err := r.DB.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.EmailVerifiedAt,
			&user.Phone,
			&user.PhoneVerifiedAt,
			&user.DisplayName,
			&user.AvatarUrl,
			&user.Status,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	// Return first user (sesuai signature, tapi idealnya kembalikan slice)
	return users[0], nil
}
