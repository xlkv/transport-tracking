-- +goose Up

CREATE TABLE IF NOT EXISTS vehicles(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    number INTEGER NOT NULL CHECK (number > 0),
    status TEXT NOT NULL CHECK (status IN ('active', 'inactive', 'maintenance')),
    driver_id BIGINT NOT NULL REFERENCES drivers(id),
    route_id BIGINT NOT NULL REFERENCES routes(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS vehicles;