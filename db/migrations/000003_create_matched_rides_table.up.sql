CREATE TABLE matched_rides (
    id SERIAL PRIMARY KEY,
    passenger_id INT REFERENCES passengers(id) ON DELETE CASCADE,
    driver_id INT REFERENCES drivers(id) ON DELETE SET NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);