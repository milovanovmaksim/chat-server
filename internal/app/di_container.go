package app

import (
	"context"
	"log"

	"github.com/milovanovmaksim/chat-server/internal/client/database"
	"github.com/milovanovmaksim/chat-server/internal/client/database/postgresql"
	"github.com/milovanovmaksim/chat-server/internal/client/database/transaction"
	"github.com/milovanovmaksim/chat-server/internal/closer"
	"github.com/milovanovmaksim/chat-server/internal/repository"
	chatRepo "github.com/milovanovmaksim/chat-server/internal/repository/chat"
	"github.com/milovanovmaksim/chat-server/internal/server"
	"github.com/milovanovmaksim/chat-server/internal/server/grpc"
	"github.com/milovanovmaksim/chat-server/internal/service"
	"github.com/milovanovmaksim/chat-server/internal/service/chat"
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
		di.chatRepository = chatRepo.NewChatRepository(di.DBClient(ctx))
	}

	return di.chatRepository
}

// ChatService возврашщает объект, удовлетворяющий интерфейсу service.ChatService.
func (di *diContainer) ChatService(ctx context.Context) service.ChatService {
	if di.chatService == nil {
		di.chatService = chat.NewChatService(di.ChatRepository(ctx), di.TxManager(ctx))
	}

	return di.chatService
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

// GRPCConfig возвращает объект, удовлетворяющий интерфейсу server.ServerConfig.
func (di *diContainer) GRPCConfig() server.ServerConfig {
	if di.grpcConfig == nil {
		cfg, err := grpc.NewGrpcConfigFromEnv()
		if err != nil {
			log.Fatalf("failed to get grpc config || error: %v", err.Error())
		}

		di.grpcConfig = cfg
	}

	return di.grpcConfig
}

// TxManager возвращает объект, удовлетвoряющий интерфейсу database.TxManager.
func (di *diContainer) TxManager(ctx context.Context) database.TxManager {
	if di.txManager == nil {
		di.txManager = transaction.NewTransactionManager(di.DBClient(ctx).DB())
	}

	return di.txManager
}
