-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_action_history (
    id BIGSERIAL NOT NULL,
    user_id BIGINT NOT NULL,
    profile_id BIGINT NOT NULL,
    action VARCHAR(10) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT user_action_history_pkey PRIMARY KEY (id),
    CONSTRAINT user_action_history_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT user_action_history_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_action_history;
-- +goose StatementEnd
