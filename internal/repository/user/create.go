package user

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/milovanovmaksim/chat-server/internal/client/database"
)

// CreateUser создает нового пользователя в БД.
func (u *userRepositoryImpl) CreateUser(ctx context.Context, userID int64) (int64, error) {
	var response int64

	queryRow := "INSERT INTO users (user_id) VALUES($1) ON CONFLICT DO NOTHING RETURNING user_id"

	query := database.Query{Name: "Create user", QueryRaw: queryRow}

	err := u.db.DB().ScanOneContext(ctx, &response, query, userID)
	if err != nil {
		// Проверка на дубликат. Ни одной строки не вернется, если пользователь уже существует.
		if errors.Is(err, pgx.ErrNoRows) {
			return response, nil
		}

		log.Printf("failed to insert user: %v", err)
		return 0, err
	}

	return response, nil
}
