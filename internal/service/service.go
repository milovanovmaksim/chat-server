package service

import "context"

type ChatService interface {
	CreateChat(ctx context.Context, request CreateChatRequest) (*CreateChatResponse, error)
	DeleteChat(ctx context.Context, request DeleteChatRequest) error
}
