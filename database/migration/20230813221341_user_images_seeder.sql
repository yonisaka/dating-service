-- +goose Up
-- +goose StatementBegin
INSERT INTO user_images (user_id, image_url, created_at, updated_at)
VALUES (1, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (2, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (3, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (4, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (5, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (6, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (7, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (8, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (9, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (10, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (11, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (12, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (13, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (14, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (15, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (16, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (17, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (18, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (19, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (20, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (21, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (22, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (23, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (24, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (25, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (26, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (27, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW()),
       (28, 'fed41671-5514-4793-a6ae-6f34484710e5', NOW(), NOW());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM user_images;
-- +goose StatementEnd
