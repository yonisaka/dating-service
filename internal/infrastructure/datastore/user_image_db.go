package datastore

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/yonisaka/dating-service/internal/entities/repository"
)

// userImage is a struct
type userImage struct {
	*BaseRepo
}

// NewUserImageRepo is a function
func NewUserImageRepo(base *BaseRepo) repository.UserImageRepo {
	return &userImage{
		BaseRepo: base,
	}
}

// FindByUserID is a method
func (r *userImage) FindByUserID(ctx context.Context, userID int64) ([]repository.UserImage, error) {
	var userImages []repository.UserImage

	query := `
		SELECT id, user_id, image_url
		FROM user_images
		WHERE user_id = $1
	`

	rows, err := r.dbMaster.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userImage repository.UserImage
		if err := rows.Scan(
			&userImage.ID,
			&userImage.UserID,
			&userImage.ImageURL,
		); err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}

		userImages = append(userImages, userImage)
	}

	if err == pgx.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return userImages, nil
}

// StoreBulk is a method
func (r *userImage) StoreBulk(ctx context.Context, images []repository.UserImage) error {
	query := `
		INSERT INTO user_images (user_id, image_url)
		VALUES
	`

	for i, image := range images {
		if i == 0 {
			query += fmt.Sprintf(" (%d, '%s')", image.UserID, image.ImageURL)
			continue
		}

		query += fmt.Sprintf(", (%d, '%s')", image.UserID, image.ImageURL)
	}

	_, err := r.dbSlave.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
