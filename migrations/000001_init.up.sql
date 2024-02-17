CREATE TABLE IF NOT EXISTS players
(
    id            SERIAL PRIMARY KEY
    email         VARCHAR(255) UNIQUE
    nickname      VARCHAR(255) UNIQUE
    password_hash VARCHAR(255)
);

CREATE INDEX IF NOT EXISTS idx_players_email ON players;