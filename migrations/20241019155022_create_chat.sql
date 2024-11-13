-- +goose Up
CREATE TABLE chats (
  id serial PRIMARY KEY,
  title TEXT not null unique,
  user_ids bigint [],
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE chats;
