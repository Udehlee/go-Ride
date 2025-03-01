CREATE TABLE matched_rides (
    id SERIAL PRIMARY KEY,
    passenger_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    driver_id UUID REFERENCES users(id) ON DELETE SET NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    CHECK (passenger_id <> driver_id)
);
