package chat

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/service/chat/model"
)

// CreateChat создает новый чат.
func (c *chatServiceImpl) CreateChat(ctx context.Context, request model.CreateChatRequest) (int64, error) {
	var chatID int64

	errValid := ValidateInputData(request)
	if errValid != nil {
		return 0, errValid
	}

	err := c.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		chatID, errTx = c.chatRepository.CreateChat(ctx, request.TitleChat)
		if errTx != nil {
			log.Printf("failed to create new chat: %v", errTx)
			return errTx
		}

		for _, userID := range request.UserIDs {
			_, errTx = c.userRepository.CreateUser(ctx, userID)
			if errTx != nil {
				log.Printf("failed to create new user: %v", errTx)
				return errTx
			}

			_, errTx = c.chatRepository.CreateChatUser(ctx, userID, chatID)
			if errTx != nil {
				log.Printf("failed to create new chat_user: %v", errTx)
				return errTx
			}
		}

		return nil
	})
	if err != nil {
		log.Printf("failed to create new chat: %v", err)
		return 0, err
	}

	return chatID, nil
}


