package repository

// CreateChatResponse ответ на запрос о создании нового чата.
type CreateChatResponse struct {
	ID int64 `db:"id"`
}

type CreateUserResponse struct {
	ID int64
}
