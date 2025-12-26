-- +goose Up
CREATE TABLE users(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    email       TEXT
);

-- +goose Down
DROP TABLE users;
