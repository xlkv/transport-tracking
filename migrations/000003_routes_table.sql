-- +goose Up
CREATE TABLE IF NOT EXISTS routes(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    from_stop_id BIGINT NOT NULL REFERENCES stops(id),
    to_stop_id BIGINT NOT NULL REFERENCES stops(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS routes;
