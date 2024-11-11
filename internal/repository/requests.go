package repository

type CreateChatRequest struct {
	UserIds []int64
	TitleChat string
}



type DeleteChatRequest struct {
	Id int64
}