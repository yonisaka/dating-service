package repository

import (
	"context"
	"time"
)

// UserActionHistory is a user action history entity.
type UserActionHistory struct {
	ID        int64
	UserID    int64
	ProfileID int64
	Action    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// UserActionHistoryRepo is a user action history repository interface.
type UserActionHistoryRepo interface {
	FindByUserID(ctx context.Context, userID int64) ([]UserActionHistory, error)
	Store(ctx context.Context, history UserActionHistory) error
}
