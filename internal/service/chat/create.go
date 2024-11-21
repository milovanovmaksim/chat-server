package chat

import (
	"context"
	"errors"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/repository"
	"github.com/milovanovmaksim/chat-server/internal/service"
)

// CreateChat создает новый чат.
func (c *chatServiceImpl) CreateChat(ctx context.Context, request service.CreateChatRequest) (*service.CreateChatResponse, error) {
	var chat *repository.CreateChatResponse
	var user *repository.CreateUserResponse

	if len(request.UserIDs) == 0 {
		return nil, errors.New("user_ids is empty")
	}

	err := c.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		chat, errTx = c.chatRepository.CreateChat(ctx, request.Into())
		if errTx != nil {
			log.Printf("failed to create new chat || error: %v", errTx)
			return errTx
		}

		for _, userID := range request.UserIDs {
			user, errTx = c.userRepository.CreateUser(ctx, repository.CreateUserRequest{UserID: userID})
			if errTx != nil {
				log.Printf("failed to create new chat || error: %v", errTx)
				return errTx
			}

			_, errTx = c.chatRepository.CreateChatUser(ctx, user.ID, chat.ID)
			if errTx != nil {
				log.Printf("failed to create new chat || error: %v", errTx)
				return errTx
			}
		}

		return nil
	})
	if err != nil {
		log.Printf("failed to create new chat || error: %v", err)
		return nil, err
	}

	return &service.CreateChatResponse{ID: chat.ID}, nil
}
