package chat

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/client/database"
	"github.com/milovanovmaksim/chat-server/internal/repository"
)

// CreateChat создает новый чат, сохраняет информацию о чате в БД.
func (c *chatRepositoryImpl) CreateChat(ctx context.Context, request repository.CreateChatRequest) (*repository.CreateChatResponse, error) {
	var response repository.CreateChatResponse

	query := database.Query{Name: "Create chat", QueryRaw: "INSERT INTO chats (title, user_ids) VALUES($1, $2) RETURNING id"}
	err := c.db.DB().ScanOneContext(ctx, response, query, request.TitleChat, request.UserIDs)
	if err != nil {
		log.Printf("failed to insert user || err: %v", err)
		return nil, err
	}

	return &response, nil
}
