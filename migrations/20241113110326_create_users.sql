-- +goose Up
CREATE TABLE users (
  user_id bigint PRIMARY KEY,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE TABLE chat_users (
  id serial PRIMARY KEY,
  chat_id bigint not null references chats (id) on delete cascade,
  user_id bigint not null references users (user_id) on delete cascade,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
    
);

-- +goose Down
DROP TABLE users;
DROP TABLE chat_users;
