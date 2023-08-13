-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX user_preferences_user_id_idx ON user_preferences (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS user_preferences_user_id_idx;
-- +goose StatementEnd
