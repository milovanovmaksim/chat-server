package service

import (
	"context"

	"github.com/milovanovmaksim/chat-server/internal/service/chat/model"
)

// ChatService интерфейс, отвечающий за бизнес логику приложения.
type ChatService interface {
	CreateChat(ctx context.Context, request model.CreateChatRequest) (int64, error)
	DeleteChat(ctx context.Context, request int64) error
}
