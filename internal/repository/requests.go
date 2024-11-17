package repository

// CreateChatRequest запрос на создание нового чата.
type CreateChatRequest struct {
	TitleChat string
	UserIDs   []int64
}

// DeleteChatRequest запрос на удаление чата.
type DeleteChatRequest struct {
	ID int64
}

// CreateUserRequest запрос на создание нового пользователя чатом.
type CreateUserRequest struct {
	userID int64
}

// DeleteUserRequest запрос на удаление пользователя из чата.
type DeleteUserRequest struct {
	userID int64
}
