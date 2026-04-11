-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY,
    password text NOT NULL,
    email text NOT NULL UNIQUE,
    first_name text NOT NULL,
    last_name text NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);

-- +goose Down
DROP TABLE users;
