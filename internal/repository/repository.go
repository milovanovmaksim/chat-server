package repository

import "context"

// ChatRepository интерфейс, определяющий набор методов CRUD для работы с БД.
type ChatRepository interface {
	CreateChat(ctx context.Context, request CreateChatRequest) (*CreateChatResponse, error)
	DeleteCaht(ctx context.Context, request DeleteChatRequest) error
}
