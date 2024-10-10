package main

import (
	"log"

	"github.com/olezhek28/microservices_course_boilerplate/cmd/server"
)

const grpcPort = 50051

func main() {
	server := server.Server{}
	err := server.Start(grpcPort)
	if err != nil {
		log.Fatalf("Failed to start server | error: %v", err)
	}
}
