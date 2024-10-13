package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/milovanovmaksim/chat-server/pkg/chat_v1"
)

// Server ...
type Server struct {
	desc.UnimplementedChatV1Server
}

// Create ...
func (s *Server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create new chat with users: %v", req.Usernames)
	return &desc.CreateResponse{Id: 1}, nil
}

// Delete ...
func (s *Server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete chat with id: %d", req.GetId())
	return &emptypb.Empty{}, nil
}

// SendMessage ...
func (s *Server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Send message with content: %v", req.Message)
	return &emptypb.Empty{}, nil
}

// Start ...
func (s *Server) Start(grpcPort int64) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	server := grpc.NewServer()
	reflection.Register(server)
	desc.RegisterChatV1Server(server, s)
	log.Printf("server listening at %v", lis.Addr())
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}
