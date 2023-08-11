package datastore

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/yonisaka/dating-service/internal/entities/repository"
)

// userRepo is a user repository.
type userRepo struct {
	*BaseRepo
}

// NewUserRepo returns a user repository.
func NewUserRepo(base *BaseRepo) repository.UserRepo {
	return &userRepo{
		BaseRepo: base,
	}
}

// FindByEmail returns a user by email.
func (r *userRepo) FindByEmail(ctx context.Context, email string) (*repository.User, error) {
	var user repository.User

	query := `
		SELECT id, first_name, last_name, email, phone, password, 
		    dob, gender, intend, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	err := r.dbSlave.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.Dob,
		&user.Gender,
		&user.Intend,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("user email: %s not found", email)
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Store stores a user.
func (r *userRepo) Store(ctx context.Context, user *repository.User) error {
	query := `
		INSERT INTO users (first_name, last_name, email, phone, password, dob, gender, intend, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
	`
	_, err := r.dbSlave.Exec(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Password,
		user.Dob,
		user.Gender,
		user.Intend,
	)

	if err != nil {
		return err
	}

	return nil
}
