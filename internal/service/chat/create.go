package chat

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/repository"
	"github.com/milovanovmaksim/chat-server/internal/service"
)

// CreateChat создает новый чат.
func (c *chatServiceImpl) CreateChat(ctx context.Context, request service.CreateChatRequest) (*service.CreateChatResponse, error) {
	chat, err := c.chatRepository.CreateChat(ctx, repository.CreateChatRequest{TitleChat: request.TitleChat, UserIds: request.UserIds})
	if err != nil {
		log.Printf("failed to create new chat || error: %v", err)
		return nil, err
	}

	return &service.CreateChatResponse{Id: chat.Id}, nil
}
