package main

import (
	"context"
	"log"

	desc "github.com/olezhek28/microservices_course_boilerplate/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
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

// Delete ...
func (s *server) Delete(_ context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	log.Printf("Delete chat with id: %d", req.GetId())
	return &emptypb.Empty{}, nil
}

// SendMessage ...
func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Send message with content: %v", req.Message)
	return &emptypb.Empty{}, nil
}
