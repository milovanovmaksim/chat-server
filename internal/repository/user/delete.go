package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/client/database"
	"github.com/milovanovmaksim/chat-server/internal/repository"
)

// DeleteUser удаляет пользователя из БД.
func (u *userRepositoryImpl) DeleteUser(ctx context.Context, request repository.DeleteUserRequest) error {
	query := database.Query{Name: "Delete chat", QueryRaw: "DELETE FROM users WHERE id = $1"}

	_, err := u.db.DB().ExecContext(ctx, query, request.UserID)
	if err != nil {
		log.Printf("failed to delete user || error: %v", err)
		return err
	}

	return nil
}
