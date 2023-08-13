-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_preferences (
    id BIGSERIAL NOT NULL,
    user_id BIGINT NOT NULL,
    min_age INT NOT NULL,
    max_age INT NOT NULL,
    use_intend BOOL NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT user_preferences_pkey PRIMARY KEY (id),
    CONSTRAINT user_preferences_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_preferences;
-- +goose StatementEnd
