package service

import (
	"database/sql"

	"github.com/milovanovmaksim/chat-server/internal/repository"
)

// CreateChatRequest запрос на создание нового чата.
type CreateChatRequest struct {
	TitleChat string
	UserIDs   []int64
}

// Into преобразует из service.CreateChatRequest в repository.CreateChatRequest.
func (c CreateChatRequest) Into() repository.CreateChatRequest {
	var titleChat sql.NullString

	if c.TitleChat == "" {
		titleChat = sql.NullString{String: "", Valid: false}
	} else {
		titleChat = sql.NullString{String: c.TitleChat, Valid: true}
	}

	return repository.CreateChatRequest{
		TitleChat: titleChat,
	}
}

// DeleteChatRequest запрос на удаление чата.
type DeleteChatRequest struct {
	ID int64
}
