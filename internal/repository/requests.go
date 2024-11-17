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

type CreateUserRequest struct {
	userID int64
}

type DeleteUserRequest struct {
	userID int64
}
