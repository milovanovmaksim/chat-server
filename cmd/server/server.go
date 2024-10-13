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

// Server - чат-сервер.
type Server struct {
	desc.UnimplementedChatV1Server
}

// CreateChat создание нового чата.
func (s *Server) CreateChat(_ context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	log.Printf("Create new chat with title: %s and user ids: %+v", req.GetTitleChat(), req.GetUserIds())
	return &desc.CreateChatResponse{Id: 1}, nil
}

// DeleteChat удаление чата.
func (s *Server) DeleteChat(_ context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	log.Printf("Delete chat with id: %d", req.GetId())
	return &emptypb.Empty{}, nil
}

// SendMessage отправление сообщения в чат.
func (s *Server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Send message with content: %s, from user with id: %d",
		req.GetMessage().GetText(), req.GetMessage().GetFrom())
	return &emptypb.Empty{}, nil
}

// Start старт чат-сервера.
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
