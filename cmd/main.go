package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"

	"github.com/milovanovmaksim/chat-server/cmd/server"
	grpc_config "github.com/milovanovmaksim/chat-server/internal/config"
	"github.com/milovanovmaksim/chat-server/internal/pgsql"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config || err: %v", err)
	}

	ctx := context.Background()

	dbConfig, err := pgsql.NewConfigFromEnv()
	if err != nil {
		log.Fatalf("failed to load config || err: %v", err)
	}

	grpcConfig, err := grpc_config.NewGrpcConfigFromEnv()
	if err != nil {
		log.Fatalf("failed to load grpc config || err: %v", err)
	}

	postgreSQL, err := pgsql.Connect(ctx, dbConfig)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL || err: %v", err)
	}

	defer postgreSQL.Close()

	server := server.NewServer(postgreSQL, grpcConfig)
	err = server.Start()
	if err != nil {
		log.Fatalf("failed to start a server || error: %v", err)
	}
}
