package service

import "context"

// ChatService интерфейс, отвечающий за бизнес логику приложения.
type ChatService interface {
	CreateChat(ctx context.Context, request CreateChatRequest) (*CreateChatResponse, error)
	DeleteChat(ctx context.Context, request DeleteChatRequest) error
}
