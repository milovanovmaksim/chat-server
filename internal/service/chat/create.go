package chat

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/repository"
	"github.com/milovanovmaksim/chat-server/internal/service"
)

// CreateChat создает новый чат.
func (c *chatServiceImpl) CreateChat(ctx context.Context, request service.CreateChatRequest) (*service.CreateChatResponse, error) {
	var chat *repository.CreateChatResponse
	var user *repository.CreateUserResponse

	err := c.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		chat, errTx = c.chatRepository.CreateChat(ctx, repository.CreateChatRequest{TitleChat: request.TitleChat})
		if errTx != nil {
			log.Printf("failed to create new chat || error: %v", errTx)
			return errTx
		}

		for _, userID := range request.UserIDs {
			ok, errTx := c.userRepository.UserExists(ctx, userID)
			if errTx != nil {
				log.Printf("failed to create new chat || error: %v", errTx)
				return errTx
			}

			if !ok {
				user, errTx = c.userRepository.CreateUser(ctx, repository.CreateUserRequest{UserID: userID})
				if errTx != nil {
					log.Printf("failed to create new chat || error: %v", errTx)
					return errTx
				}
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
