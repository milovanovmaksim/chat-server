package repository

// CreateChatRequest запрос на создание нового чата.
type CreateChatRequest struct {
	TitleChat string
	UserIds   []int64
}

// DeleteChatRequest запрос на удаление чата.
type DeleteChatRequest struct {
	Id int64
}
