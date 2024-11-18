package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/client/database"
	"github.com/milovanovmaksim/chat-server/internal/repository"
)

// CreateUser создает нового пользователя в БД.
func (u *userRepositoryImpl) CreateUser(ctx context.Context, request repository.CreateUserRequest) (*repository.CreateUserResponse, error) {
	var response repository.CreateUserResponse

	query := database.Query{Name: "Create user", QueryRaw: "INSERT INTO users (user_id) VALUES($1) RETURNING user_id"}
	err := u.db.DB().ScanOneContext(ctx, &response, query, request.UserID)
	if err != nil {
		log.Printf("failed to insert user || err: %v", err)
		return nil, err
	}

	return &response, nil
}
