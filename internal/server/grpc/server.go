package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/milovanovmaksim/chat-server/internal/server"
	"github.com/milovanovmaksim/chat-server/internal/service"
	desc "github.com/milovanovmaksim/chat-server/pkg/chat_v1"
)

// Server - чат-сервер.
type Server struct {
	desc.UnimplementedChatV1Server
	grpcConfig server.ServerConfig
	grpcServer *grpc.Server
	service    service.ChatService
}

// NewServer создает новый Server объект.
func NewServer(grpcConfig server.ServerConfig, service service.ChatService) Server {
	return Server{desc.UnimplementedChatV1Server{}, grpcConfig, nil, service}
}

// CreateChat создание нового чата.
func (s *Server) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	chat, err := s.service.CreateChat(ctx, service.CreateChatRequest{
		TitleChat: req.TitleChat,
		UserIds:   req.UserIds,
	})
	if err != nil {
		log.Printf("failed to create new chat || error: %v", err)
		return nil, err
	}

	return &desc.CreateChatResponse{
		Id: chat.Id,
	}, nil
}

// DeleteChat удаление чата.
func (s *Server) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// SendMessage отправление сообщения в чат.
func (s *Server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Send message with content: %s, from user with id: %d",
		req.GetMessage().GetText(), req.GetMessage().GetFrom())
	return &emptypb.Empty{}, nil
}

// Start старт чат-сервера.
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	s.grpcServer = grpc.NewServer()
	reflection.Register(s.grpcServer)
	desc.RegisterChatV1Server(s.grpcServer, s)
	log.Printf("server listening at %v", lis.Addr())

	if err = s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}

// Stop остановка сервера.
func (s *Server) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.Stop()
	}
}
