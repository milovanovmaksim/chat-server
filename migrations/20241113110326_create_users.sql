-- +goose Up
CREATE TABLE users (
  id serial PRIMARY KEY,
  user_id bigint not null unique,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE TABLE chat_users (
  id bigint PRIMARY KEY,
  chat_id bigint not null references chats (id) on delete cascade,
  user_id bigint not null references users (user_id) on delete cascade,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
    
);

-- +goose Down
DROP TABLE chats;
DROP TABLE chat_users;
