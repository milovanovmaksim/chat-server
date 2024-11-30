package chat

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/client/database"
)

// CreateChat создает новый чат в БД.
func (c *chatRepositoryImpl) CreateChat(ctx context.Context, chatTitle string) (int64, error) {
	var chatID int64

	queryRow := "INSERT INTO chats (title) VALUES($1) RETURNING id"

	query := database.Query{Name: "Create chat", QueryRaw: queryRow}

	err := c.db.DB().ScanOneContext(ctx, &chatID, query, chatTitle)
	if err != nil {
		log.Printf("failed to insert chat: %v", err)
		return 0, err
	}

	return chatID, nil
}

// CreateChatUser создает запись в таблицу "user_chats".
func (c *chatRepositoryImpl) CreateChatUser(ctx context.Context, userID int64, chatID int64) (int64, error) {
	var id int64

	queryRow := "INSERT INTO chat_users (user_id, chat_id) VALUES($1, $2) RETURNING id"

	query := database.Query{Name: "Create chat user", QueryRaw: queryRow}

	err := c.db.DB().ScanOneContext(ctx, &id, query, userID, chatID)
	if err != nil {
		log.Printf("failed to create chat_user: %v", err)
		return 0, err
	}

	return id, nil
}
