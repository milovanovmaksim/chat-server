-- +goose Up
CREATE TABLE chats (
  id serial PRIMARY KEY,
  title varchar(64) not null unique,
  user_ids bigint [],
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE users;
