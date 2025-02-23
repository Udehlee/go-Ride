CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    user_role VARCHAR(20) CHECK (role IN ('passenger', 'driver')) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
