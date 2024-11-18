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

func NewUserRepository(db database.Client) repository.UserRepository {
	return &userRepositoryImpl{db}
}

func (u *userRepositoryImpl) UserExists(ctx context.Context, request int64) (bool, error) {
	var count int64

	query := database.Query{Name: "Does user exist?", QueryRaw: "SELECT COUNT(user_id) FROM users WHERE user_id=$1"}
	err := u.db.DB().ScanOneContext(ctx, &count, query, request)
	if err != nil {
		log.Printf("failed to count users || error: %v", err)
		return false, err
	}

	return count == 1, nil
}
