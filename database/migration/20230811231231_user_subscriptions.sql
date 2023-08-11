-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_subscriptions
(
    id BIGSERIAL NOT NULL,
    user_id BIGINT NOT NULL,
    subscription_code VARCHAR(50) NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT user_subscriptions_pkey PRIMARY KEY (id),
    CONSTRAINT user_subscriptions_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_subscriptions;
-- +goose StatementEnd
