package chat

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/client/database"
)

// DeleteCaht удаляет чат из БД.
func (c *chatRepositoryImpl) DeleteCaht(ctx context.Context, chatID int64) error {
	queryRow := "DELETE FROM CHATS WHERE id = $1"

	query := database.Query{Name: "Delete chat", QueryRaw: queryRow}

	_, err := c.db.DB().ExecContext(ctx, query, chatID)
	if err != nil {
		log.Printf("failed to delete chat: %v", err)
		return err
	}

	return nil
}
