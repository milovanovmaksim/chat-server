package repository

import "context"

// ChatRepository интерфейс, определяющий набор методов CRUD для работы с БД.
type ChatRepository interface {
	CreateChat(ctx context.Context, request CreateChatRequest) (*CreateChatResponse, error)
	DeleteCaht(ctx context.Context, request DeleteChatRequest) error
}

// UserRepository интерфейс, определяющий набор методов для работы с таблицей "users" базы данных.
type UserRepository interface {
	CreateUser(ctx context.Context, request CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(ctx context.Context, request DeleteUserRequest) error
}
