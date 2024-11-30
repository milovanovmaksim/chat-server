package chat

import "github.com/milovanovmaksim/chat-server/internal/service/chat/model"

func validateUserIDs(userIDs []int64) error {
	if len(userIDs) == 0 {
		return ValidationError{"field 'user_ids' is empty"}
	}

	return nil
}

func validateChatTitle(chatTitle string) error {
	if chatTitle == "" {
		return ValidationError{"field 'title' is empty"}
	}

	return nil
}

// ValidateInputData валидирует входные данные при создании нового чата.
func ValidateInputData(request model.CreateChatRequest) error {
	err := validateUserIDs(request.UserIDs)
	if err != nil {
		return err
	}

	err = validateChatTitle(request.TitleChat)
	if err != nil {
		return err
	}

	return nil
}
