package main

import (
	"context"
	"log"

	desc "github.com/olezhek28/microservices_course_boilerplate/pkg/chat_v1"
)

// Server ...
type server struct {
	desc.UnimplementedChatServer
}

// Create ...
func (s *server) Create(_ context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	log.Printf("Create new chat with users: %v", req.Usernames)
	return &desc.CreateChatResponse{Id: 1}, nil
}
