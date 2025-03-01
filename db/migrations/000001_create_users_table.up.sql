CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    pass_word TEXT NOT NULL,
    role VARCHAR(20) CHECK (role IN ('driver', 'passenger')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
