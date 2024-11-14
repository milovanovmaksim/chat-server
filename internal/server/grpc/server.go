package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/milovanovmaksim/chat-server/internal/closer"
	"github.com/milovanovmaksim/chat-server/internal/server"
	"github.com/milovanovmaksim/chat-server/internal/service"
	desc "github.com/milovanovmaksim/chat-server/pkg/chat_v1"
)

// Server - чат-сервер.
type Server struct {
	desc.UnimplementedChatV1Server
	grpcConfig server.Config
	grpcServer *grpc.Server
	service    service.ChatService
}

// NewServer создает новый Server объект.
func NewServer(grpcConfig server.Config, service service.ChatService) Server {
	return Server{desc.UnimplementedChatV1Server{}, grpcConfig, nil, service}
}

// CreateChat создание нового чата.
func (s *Server) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	chat, err := s.service.CreateChat(ctx, service.CreateChatRequest{
		TitleChat: req.TitleChat,
		UserIDs:   req.UserIds,
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
	err := s.service.DeleteChat(ctx, service.DeleteChatRequest{Id: req.Id})
	if err != nil {
		log.Printf("failed to delete the chat || error: %v", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// SendMessage отправление сообщения в чат.
func (s *Server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Send message with content: %s, from user with id: %d",
		req.GetMessage().GetText(), req.GetMessage().GetFrom())
	return &emptypb.Empty{}, nil
}

// Start cтарт чат-сервера.
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.grpcConfig.Address())
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}

	closer.Add(lis.Close)

	s.grpcServer = grpc.NewServer()

	reflection.Register(s.grpcServer)
	desc.RegisterChatV1Server(s.grpcServer, s)
	log.Printf("server listening at %v", lis.Addr())

	if err = s.grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
		return err
	}

	closer.Add(func() error {
		s.grpcServer.Stop()
		return nil
	})

	return nil
}
