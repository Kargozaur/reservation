-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY,
    password text NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

-- +goose Down
DROP TABLE users;
