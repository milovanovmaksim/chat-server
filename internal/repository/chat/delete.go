package chat

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/client/database"
	"github.com/milovanovmaksim/chat-server/internal/repository"
)

// DeleteCaht удаляет чат из БД.
func (c *chatRepositoryImpl) DeleteCaht(ctx context.Context, request repository.DeleteChatRequest) error {
	query := database.Query{Name: "Delete chat", QueryRaw: "DELETE FROM CHATS WHERE id = $1"}

	_, err := c.db.DB().ExecContext(ctx, query, request.ID)
	if err != nil {
		log.Printf("failed to delete user || error: %v", err)
		return err
	}

	return nil
}
