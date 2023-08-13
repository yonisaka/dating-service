package repository

import "context"

// UserImage is a user image entity.
type UserImage struct {
	ID       int64
	UserID   int64
	ImageURL string
}

// UserImageRepo is a user image repository interface.
type UserImageRepo interface {
	FindByUserID(ctx context.Context, userID int64) ([]UserImage, error)
	StoreBulk(ctx context.Context, images []UserImage) error
}
