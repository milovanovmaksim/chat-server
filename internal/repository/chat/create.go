package chat

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/client/database"
	"github.com/milovanovmaksim/chat-server/internal/repository"
)

// CreateChat создает новый чат в БД.
func (c *chatRepositoryImpl) CreateChat(ctx context.Context, request repository.CreateChatRequest) (*repository.CreateChatResponse, error) {
	var response repository.CreateChatResponse

	query := database.Query{Name: "Create chat", QueryRaw: "INSERT INTO chats (title) VALUES($1) RETURNING id"}

	err := c.db.DB().ScanOneContext(ctx, &response, query, request.TitleChat)
	if err != nil {
		log.Printf("failed to insert chat || err: %v", err)
		return nil, err
	}

	return &response, nil
}

// CreateChatUser создает запись в таблицу "user_chats".
func (c *chatRepositoryImpl) CreateChatUser(ctx context.Context, userID int64, chatID int64) (int64, error) {
	var id int64

	query := database.Query{Name: "Create chat user", QueryRaw: "INSERT INTO chat_users (user_id, chat_id) VALUES($1, $2) RETURNING id"}

	err := c.db.DB().ScanOneContext(ctx, &id, query, userID, chatID)
	if err != nil {
		log.Printf("failed to create chat_user || error: %v", err)
		return 0, err
	}

	return id, nil
}
