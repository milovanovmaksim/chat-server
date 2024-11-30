package chat

import (
	"github.com/milovanovmaksim/chat-server/internal/client/database"
	"github.com/milovanovmaksim/chat-server/internal/repository"
)

type chatRepositoryImpl struct {
	db database.Client
}

// NewChatRepository создает новый объект, удовлетворяющий интерфейсу repository.ChatRepository.
func NewChatRepository(db database.Client) repository.ChatRepository {
	return &chatRepositoryImpl{db}
}
