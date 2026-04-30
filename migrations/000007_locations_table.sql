-- +goose Up
CREATE TABLE IF NOT EXISTS locations(
    id BIGSERIAL PRIMARY KEY,
    vehicle_id BIGINT NOT NULL REFERENCES vehicles(id),
    trip_id BIGINT NOT NULL REFERENCES trips(id),
    lat DECIMAL(9,6) NOT NULL,
    lng DECIMAL(9,6) NOT NULL,
    recorded_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_vehicle_id ON locations(vehicle_id);
CREATE INDEX IF NOT EXISTS idx_trip_id ON locations(trip_id);
CREATE INDEX IF NOT EXISTS idx_recorded_at ON locations(recorded_at);

-- +goose Down
DROP TABLE IF EXISTS locations; 