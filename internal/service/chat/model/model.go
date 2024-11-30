package model

// CreateChatRequest запрос на создание нового чата.
type CreateChatRequest struct {
	TitleChat string
	UserIDs   []int64
}
