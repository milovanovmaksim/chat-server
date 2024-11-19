package chat

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/repository"
	"github.com/milovanovmaksim/chat-server/internal/service"
)

// CreateChat создает новый чат.
func (c *chatServiceImpl) CreateChat(ctx context.Context, request service.CreateChatRequest) (*service.CreateChatResponse, error) {
	var chat *repository.CreateChatResponse
	var createChatRequest repository.CreateChatRequest

	if len(request.UserIDs) == 0 {
		return nil, errors.New("user_ids is empty")
	}

	err := c.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		if request.TitleChat == "" {
			createChatRequest = repository.CreateChatRequest{TitleChat: sql.NullString{Valid: false}}
		} else {
			createChatRequest = repository.CreateChatRequest{TitleChat: sql.NullString{String: request.TitleChat, Valid: true}}
		}

		chat, errTx = c.chatRepository.CreateChat(ctx, createChatRequest)
		if errTx != nil {
			log.Printf("failed to create new chat || error: %v", errTx)
			return errTx
		}

		for _, userID := range request.UserIDs {
			ok, errTx := c.userRepository.UserExists(ctx, userID)
			if errTx != nil {
				log.Printf("failed to get user || error: %v", errTx)
				return errTx
			}

			if !ok {
				_, errTx = c.userRepository.CreateUser(ctx, repository.CreateUserRequest{UserID: userID})
				if errTx != nil {
					log.Printf("failed to create new chat || error: %v", errTx)
					return errTx
				}
			}

			_, errTx = c.chatRepository.CreateChatUser(ctx, userID, chat.ID)
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
