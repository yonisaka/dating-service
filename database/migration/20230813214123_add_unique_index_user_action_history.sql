-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX user_action_history_user_id_idx ON user_action_history (user_id, profile_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS user_action_history_user_id_idx;
-- +goose StatementEnd
