package repository

import "context"



type ChatRepository interface {
	CreateChat(ctx context.Context, request CreateChatRequest) (CreateChatResponse, error)
	DeleteCaht(ctx context.Context, request DeleteChatRequest) error
}
