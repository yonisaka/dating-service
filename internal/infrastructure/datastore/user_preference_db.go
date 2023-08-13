package datastore

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/yonisaka/dating-service/internal/entities/repository"
)

// userPreferenceRepo is a struct
type userPreferenceRepo struct {
	*BaseRepo
}

// NewUserPreferenceRepo is a function
func NewUserPreferenceRepo(base *BaseRepo) repository.UserPreferenceRepo {
	return &userPreferenceRepo{
		BaseRepo: base,
	}
}

// FindByUserID is a method
func (r *userPreferenceRepo) FindByUserID(ctx context.Context, userID int64) (*repository.UserPreference, error) {
	var userPreference repository.UserPreference

	query := `
		SELECT id, user_id, min_age, max_age, use_intend
		FROM user_preferences
		WHERE user_id = $1
	`

	err := r.dbMaster.QueryRow(ctx, query, userID).Scan(
		&userPreference.ID,
		&userPreference.UserID,
		&userPreference.MinAge,
		&userPreference.MaxAge,
		&userPreference.UseIntend,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &userPreference, nil
}
