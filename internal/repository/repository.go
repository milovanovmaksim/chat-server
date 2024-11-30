package repository

import "context"

// ChatRepository интерфейс, определяющий набор методов для работы с таблицей "chats" базы данных.
type ChatRepository interface {
	CreateChat(ctx context.Context, chatTitle string) (int64, error)
	DeleteCaht(ctx context.Context, request int64) error
	CreateChatUser(ctx context.Context, userID int64, chatID int64) (int64, error)
}

// UserRepository интерфейс, определяющий набор методов для работы с таблицей "users" базы данных.
type UserRepository interface {
	CreateUser(ctx context.Context, request int64) (int64, error)
	DeleteUser(ctx context.Context, request int64) error
	UserExists(ctx context.Context, request int64) (bool, error)
}
