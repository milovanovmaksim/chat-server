package main

import (
	"log"

	"github.com/milovanovmaksim/chat-server/cmd/server"
)

const grpcPort = 50051

func main() {
	server := server.Server{}
	err := server.Start(grpcPort)
	if err != nil {
		log.Fatalf("failed to start server | error: %v", err)
	}
}
