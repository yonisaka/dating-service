-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_images
(
    id BIGSERIAL NOT NULL,
    user_id BIGINT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT user_images_pkey PRIMARY KEY (id),
    CONSTRAINT user_images_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_images;
-- +goose StatementEnd
