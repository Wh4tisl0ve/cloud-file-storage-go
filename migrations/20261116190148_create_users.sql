-- +goose Up
CREATE TABLE
    users (
        id UUID PRIMARY KEY,
        username VARCHAR(15) NOT NULL UNIQUE,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
    );

-- +goose Down
DROP TABLE users;
