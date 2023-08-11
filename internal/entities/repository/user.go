package repository

import (
	"context"
	"time"
)

// User is a user entity.
type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
	Dob       *time.Time
	Gender    string
	Intend    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// UserRepo is a user repository interface.
type UserRepo interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	Store(ctx context.Context, user *User) error
}
