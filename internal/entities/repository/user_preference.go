package repository

import "context"

// UserPreference is a user preference entity.
type UserPreference struct {
	ID        int64
	UserID    int64
	MinAge    int64
	MaxAge    int64
	UseIntend bool
	Intend    string
	Gender    string
}

// UserPreferenceRepo is a user preference repository interface.
type UserPreferenceRepo interface {
	FindByUserID(ctx context.Context, userID int64) (*UserPreference, error)
	Update(ctx context.Context, userPreference UserPreference) error
}
