package repository

import (
	"context"
	"time"
)

//go:generate rm -f ./user_subscription_mock.go
//go:generate mockgen -destination user_subscription_mock.go -package repository -mock_names UserSubscriptionRepo=GoMockUserSubscriptionRepo -source user_subscription.go

// UserSubscription is a user subscription entity.
type UserSubscription struct {
	ID               int64
	UserID           int64
	SubscriptionCode string
	ExpiredAt        *time.Time
}

// UserSubscriptionRepo is a user subscription repository interface.
type UserSubscriptionRepo interface {
	FindByUserID(ctx context.Context, userID int64) (*UserSubscription, error)
	Store(ctx context.Context, userSubscription UserSubscription) error
}
