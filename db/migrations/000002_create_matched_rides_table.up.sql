CREATE TABLE matched_rides (
    matched_ride_id SERIAL PRIMARY KEY,
    driver_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    passenger_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    ride_status VARCHAR(20) CHECK (status IN ('pending', 'matched', 'completed')) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
