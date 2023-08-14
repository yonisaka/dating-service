package repository

import (
	"context"
	"time"
)

//go:generate rm -f ./user_action_history_mock.go
//go:generate mockgen -destination user_action_history_mock.go -package repository -mock_names UserActionHistoryRepo=GoMockUserActionHistoryRepo -source user_action_history.go

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
