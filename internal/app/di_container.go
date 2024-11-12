package app

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/client/database"
	"github.com/milovanovmaksim/chat-server/internal/client/database/postgresql"
	"github.com/milovanovmaksim/chat-server/internal/closer"
	"github.com/milovanovmaksim/chat-server/internal/repository"
	"github.com/milovanovmaksim/chat-server/internal/repository/chat"
	"github.com/milovanovmaksim/chat-server/internal/server"
	"github.com/milovanovmaksim/chat-server/internal/service"
)

type diContainer struct {
	chatRepository repository.ChatRepository
	chatService    service.ChatService
	dbClient       database.Client
	pgConfig       database.DBConfig
	grpcConfig     server.ServerConfig
	txManager      database.TxManager
}

func newDiContainer() diContainer {
	return diContainer{}
}

// ChatRepository возвращает объект, удовлетворяющий интерфейсу repository.ChatRepository
func (di *diContainer) ChatRepository(ctx context.Context) repository.ChatRepository {
	if di.chatRepository == nil {
		di.chatRepository = chat.NewChatRepository(di.DBClient(ctx))
	}

	return di.chatRepository
}

// DBClient возвращает объект, удовлетвoряющий интерфейсу database.Client.
func (di *diContainer) DBClient(ctx context.Context) database.Client {
	if di.dbClient == nil {
		pg, err := postgresql.Connect(ctx, di.DBConfig())
		if err != nil {
			log.Fatalf("failed to connect to PostgreSQL server")
		}

		dbClient := postgresql.NewClient(pg)
		di.dbClient = dbClient

		closer.Add(func() error {
			di.dbClient.Close()
			return nil
		})
	}

	return di.dbClient
}

// DBConfig возвращает объект, удовлетворяющий интерфейсу database.DBConfig.
func (di *diContainer) DBConfig() database.DBConfig {
	if di.pgConfig == nil {
		config, err := postgresql.NewConfigFromEnv()
		if err != nil {
			log.Fatalf("failed to get DB config || error: %v", err.Error())
		}

		di.pgConfig = config
	}

	return di.pgConfig
}
