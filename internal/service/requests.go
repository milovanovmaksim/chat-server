package service

// CreateChatRequest запрос на создание нового чата.
type CreateChatRequest struct {
	TitleChat string
	UserIDs   []int64
}

// DeleteChatRequest запрос на удаление чата.
type DeleteChatRequest struct {
	Id int64
}
