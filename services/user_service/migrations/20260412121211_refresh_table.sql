-- +goose Up
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL,
    token text NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose Down
DROP TABLE IF EXISTS refresh_tokens;
