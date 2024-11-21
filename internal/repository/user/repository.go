package user

import (
	"github.com/milovanovmaksim/chat-server/internal/client/database"
	"github.com/milovanovmaksim/chat-server/internal/repository"
)

type userRepositoryImpl struct {
	db database.Client
}

// NewUserRepository возвращает новый объект, который удовлетворяет интерфейсу repository.UserRepository.
func NewUserRepository(db database.Client) repository.UserRepository {
	return &userRepositoryImpl{db}
}
