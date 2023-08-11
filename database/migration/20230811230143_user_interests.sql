-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_interests
(
    id BIGSERIAL NOT NULL,
    user_id BIGINT NOT NULL,
    tags JSONB,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT user_interests_pkey PRIMARY KEY (id),
    CONSTRAINT user_interests_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_interests;
-- +goose StatementEnd
