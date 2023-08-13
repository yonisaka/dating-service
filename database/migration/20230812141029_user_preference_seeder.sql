-- +goose Up
-- +goose StatementBegin
INSERT INTO user_preferences (user_id, min_age, max_age, use_intend, created_at, updated_at)
VALUES (1, 18, 30, false, NOW(), NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM user_preferences WHERE user_id = 1;
-- +goose StatementEnd
