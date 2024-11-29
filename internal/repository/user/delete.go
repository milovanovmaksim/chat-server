package user

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/client/database"
)

// DeleteUser удаляет пользователя из БД.
func (u *userRepositoryImpl) DeleteUser(ctx context.Context, userID int64) error {
	queryRow := "DELETE FROM users WHERE id = $1"

	query := database.Query{Name: "Delete chat", QueryRaw: queryRow}

	_, err := u.db.DB().ExecContext(ctx, query, userID)
	if err != nil {
		log.Printf("failed to delete user: %v", err)
		return err
	}

	return nil
}
