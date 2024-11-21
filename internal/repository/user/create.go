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

	query := database.Query{Name: "Create user", QueryRaw: "INSERT INTO users (user_id) VALUES($1) ON CONFLICT DO NOTHING RETURNING user_id"}
	row, err := u.db.DB().QueryContext(ctx, query, request.UserID)
	if err != nil {
		log.Printf("failed to insert user || err: %v", err)
		return nil, err
	}

	row.Scan(&response)

	if response.ID == 0 {
		response.ID = request.UserID
	}

	return &response, nil
}
