package repository

// CreateChatRequest запрос на создание нового чата.
type CreateChatRequest struct {
	TitleChat string
}

// DeleteChatRequest запрос на удаление чата.
type DeleteChatRequest struct {
	ID int64
}

// CreateUserRequest запрос на создание нового пользователя в БД.
type CreateUserRequest struct {
	UserID int64
}

// DeleteUserRequest запрос на удаление пользователя из БД.
type DeleteUserRequest struct {
	UserID int64
}
