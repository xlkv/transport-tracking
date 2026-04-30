-- +goose Up
CREATE TABLE IF NOT EXISTS drivers (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS drivers;