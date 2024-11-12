package chat

import (
	"github.com/milovanovmaksim/chat-server/internal/client/database"
	"github.com/milovanovmaksim/chat-server/internal/repository"
	"github.com/milovanovmaksim/chat-server/internal/service"
)

type chatServiceImpl struct {
	chatRepository repository.ChatRepository
	txManager      database.TxManager
}

// NewChatService создает новый объект, удовлетворяющий интерфейсу service.ChatService.
func NewChatService(chatRepository repository.ChatRepository, txManager database.TxManager) service.ChatService {
	return &chatServiceImpl{chatRepository, txManager}
}
