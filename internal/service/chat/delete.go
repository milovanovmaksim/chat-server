package chat

import (
	"context"
	"log"
)

// DeleteChat удаляет чат.
func (c *chatServiceImpl) DeleteChat(ctx context.Context, request int64) error {
	err := c.chatRepository.DeleteCaht(ctx, request)
	if err != nil {
		log.Printf("failed to delete chat: %v", err)
		return err
	}

	return nil
}
