package repository

// CreateChatResponse ответ на запрос о создании нового чата.
type CreateChatResponse struct {
	ID int64 `db:"id"`
}

// CreateUserResponse ответ на запрос о создании нового пользователя.
type CreateUserResponse struct {
	ID int64 `db:"user_id"`
}
