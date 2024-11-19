package repository

import "database/sql"

// CreateChatRequest запрос на создание нового чата.
type CreateChatRequest struct {
	TitleChat sql.NullString
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
