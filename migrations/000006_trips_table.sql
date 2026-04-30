-- +goose Up 
CREATE TABLE IF NOT EXISTS trips(
    id BIGSERIAL PRIMARY KEY,
    vehicle_id BIGINT NOT NULL REFERENCES vehicles(id),
    route_id BIGINT NOT NULL REFERENCES routes(id),
    started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    ended_at TIMESTAMPTZ,
    status TEXT NOT NULL CHECK (status in ('pending', 'started', 'ended')),
    created_at TIMESTAMPTZ NOT NUlL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS trips;