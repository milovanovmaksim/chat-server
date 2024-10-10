package main

import (
	"github.com/olezhek28/microservices_course_boilerplate/cmd/server"
)

const grpcPort = 50051

func main() {
	server := server.Server{}
	server.Start(grpcPort)
}
