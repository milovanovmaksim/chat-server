package user

import (
	"context"
	"log"

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

// UserExists проверяет существует ли пользователь с таким id в БД.
func (u *userRepositoryImpl) UserExists(ctx context.Context, request int64) (bool, error) {
	var count int64

	queryRow := "SELECT COUNT(user_id) FROM users WHERE user_id=$1"

	query := database.Query{Name: "Does user exist?", QueryRaw: queryRow}

	err := u.db.DB().ScanOneContext(ctx, &count, query, request)
	if err != nil {
		log.Printf("failed to count users: %v", err)
		return false, err
	}

	return count == 1, nil
}
