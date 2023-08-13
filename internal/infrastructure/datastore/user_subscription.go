package datastore

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/yonisaka/dating-service/internal/entities/repository"
)

// userSubscriptionRepo is a struct
type userSubscriptionRepo struct {
	*BaseRepo
}

// NewUserSubscriptionRepo is a function
func NewUserSubscriptionRepo(base *BaseRepo) repository.UserSubscriptionRepo {
	return &userSubscriptionRepo{
		BaseRepo: base,
	}
}

// FindByUserID is a method
func (r *userSubscriptionRepo) FindByUserID(ctx context.Context, userID int64) (*repository.UserSubscription, error) {
	var userSubscription repository.UserSubscription

	query := `
		SELECT id, user_id, subscription_code, expired_at
		FROM user_subscriptions
		WHERE user_id = $1 AND expired_at > NOW()
	`

	err := r.dbMaster.QueryRow(ctx, query, userID).Scan(
		&userSubscription.ID,
		&userSubscription.UserID,
		&userSubscription.SubscriptionCode,
		&userSubscription.ExpiredAt,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &userSubscription, nil
}

// Store is a method
func (r *userSubscriptionRepo) Store(ctx context.Context, userSubscription repository.UserSubscription) error {
	query := `
		INSERT INTO user_subscriptions (user_id, subscription_code, expired_at)
		VALUES ($1, $2, $3)
	`

	_, err := r.dbMaster.Exec(ctx, query, userSubscription.UserID, userSubscription.SubscriptionCode, userSubscription.ExpiredAt)

	if err != nil {
		return err
	}

	return nil
}
