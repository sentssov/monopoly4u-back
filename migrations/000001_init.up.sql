CREATE TABLE IF NOT EXISTS players
(
    id            VARCHAR(255) PRIMARY KEY,
    email         VARCHAR(255) UNIQUE,
    nickname      VARCHAR(255) UNIQUE,
    password_hash VARCHAR(255),
    created_at    DATE,
    updated_at    DATE
);