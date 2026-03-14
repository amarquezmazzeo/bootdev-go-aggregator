-- +goose Up
CREATE TABLE feeds (
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
