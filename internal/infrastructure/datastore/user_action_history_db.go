package datastore

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/yonisaka/dating-service/internal/entities/repository"
)

// userActionHistory is a struct
type userActionHistory struct {
	*BaseRepo
}

// NewUserActionHistoryRepo is a function
func NewUserActionHistoryRepo(base *BaseRepo) repository.UserActionHistoryRepo {
	return &userActionHistory{
		BaseRepo: base,
	}
}

// FindByUserID is a method
func (r *userActionHistory) FindByUserID(ctx context.Context, userID int64) ([]repository.UserActionHistory, error) {
	var userActionHistories []repository.UserActionHistory

	query := `
		SELECT id, user_id, profile_id, action, created_at, updated_at
		FROM user_action_history
		WHERE user_id = $1
	`

	rows, err := r.dbMaster.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userActionHistory repository.UserActionHistory
		if err := rows.Scan(
			&userActionHistory.ID,
			&userActionHistory.UserID,
			&userActionHistory.ProfileID,
			&userActionHistory.Action,
			&userActionHistory.CreatedAt,
			&userActionHistory.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		userActionHistories = append(userActionHistories, userActionHistory)
	}

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return userActionHistories, nil
}

// Store is a method
func (r *userActionHistory) Store(ctx context.Context, history repository.UserActionHistory) error {
	query := `
		INSERT INTO user_action_history (user_id, profile_id, action, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		ON CONFLICT (user_id, profile_id) 
		DO UPDATE SET action = $3, updated_at = NOW()
	`

	_, err := r.dbSlave.Exec(ctx, query, history.UserID, history.ProfileID, history.Action)
	if err != nil {
		return err
	}

	return nil
}
