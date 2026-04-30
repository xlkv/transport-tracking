-- +goose Up
CREATE TABLE IF NOT EXISTS route_stops(
    id BIGSERIAL PRIMARY KEY,
    route_id BIGINT NOT NULL REFERENCES routes(id),
    stop_id BIGINT NOT NULL REFERENCES stops(id),
    stop_order INTEGER NOT NULL CHECK (stop_order > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS route_stops;