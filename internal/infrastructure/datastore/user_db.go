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

// Find returns a list of users.
func (r *userRepo) Find(ctx context.Context, preferences ...repository.UserPreference) ([]*repository.User, error) {
	var (
		userPreference repository.UserPreference
		users          []*repository.User
	)

	query := `
		SELECT id, first_name, last_name, dob, gender, intend
		FROM users
	`

	if len(preferences) > 0 {
		userPreference = preferences[0]
	}

	if userPreference.Gender != "" {
		query += fmt.Sprintf(" WHERE gender = '%s'", userPreference.Gender)
	}

	if userPreference.MinAge != 0 && userPreference.MaxAge != 0 {
		query += fmt.Sprintf(
			" AND dob BETWEEN NOW() - INTERVAL '%d years' AND NOW() - INTERVAL '%d years'",
			userPreference.MaxAge,
			userPreference.MinAge,
		)
	}

	if userPreference.UseIntend {
		query += fmt.Sprintf(" AND intend = '%s'", userPreference.Intend)
	}

	rows, err := r.dbMaster.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user repository.User
		if err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Dob,
			&user.Gender,
			&user.Intend,
		); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		users = append(users, &user)
	}

	return users, nil
}

// FindByID returns a user by id.
func (r *userRepo) FindByID(ctx context.Context, id int64) (*repository.User, error) {
	query := `
		SELECT id, first_name, last_name, email, phone, password, 
		    dob, gender, intend, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	return r.findOne(ctx, query, id)
}

// FindByEmail returns a user by email.
func (r *userRepo) FindByEmail(ctx context.Context, email string) (*repository.User, error) {
	query := `
		SELECT id, first_name, last_name, email, phone, password, 
		    dob, gender, intend, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	return r.findOne(ctx, query, email)
}

func (r *userRepo) findOne(ctx context.Context, query string, param any) (*repository.User, error) {
	var user repository.User
	err := r.dbMaster.QueryRow(ctx, query, param).Scan(
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
		return nil, fmt.Errorf("user not found")
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
