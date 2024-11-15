package chat

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/repository"
	"github.com/milovanovmaksim/chat-server/internal/service"
)

// DeleteChat удаляет чат.
func (c *chatServiceImpl) DeleteChat(ctx context.Context, request service.DeleteChatRequest) error {
	err := c.chatRepository.DeleteCaht(ctx, repository.DeleteChatRequest{ID: request.ID})
	if err != nil {
		log.Printf("failed to delete chat || error: %v", err)
		return err
	}

	return nil
}
