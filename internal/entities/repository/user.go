package repository

import (
	"context"
	"time"

	"github.com/yonisaka/dating-service/internal/consts"
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

func (u *User) Age() int {
	age := time.Now().Year() - u.Dob.Year()
	if time.Now().Month() < u.Dob.Month() {
		age--
	}

	return age
}

func (u *User) OppositeGender() string {
	if u.Gender == consts.Man {
		return consts.Woman
	}

	return consts.Man
}

// UserRepo is a user repository interface.
type UserRepo interface {
	Find(ctx context.Context, preferences ...UserPreference) ([]*User, error)
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Store(ctx context.Context, user *User) error
}
