package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	grpcConfig "github.com/milovanovmaksim/chat-server/internal/config"
	"github.com/milovanovmaksim/chat-server/internal/pgsql"
	desc "github.com/milovanovmaksim/chat-server/pkg/chat_v1"
)

// Server - чат-сервер.
type Server struct {
	pgSQL      *pgsql.PostgreSQL
	grpcConfig *grpcConfig.GrpcConfig
	desc.UnimplementedChatV1Server
}

func NewServer(pgSQL *pgsql.PostgreSQL, grpcConfig *grpcConfig.GrpcConfig) Server {
	return Server{pgSQL, grpcConfig, desc.UnimplementedChatV1Server{}}
}

// CreateChat создание нового чата.
func (s *Server) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	var id int64

	pool := s.pgSQL.GetPool()

	err := pool.QueryRow(ctx, "INSERT INTO chats (title, user_ids) VALUES($1, $2) RETURNING id", req.TitleChat, req.UserIds).Scan(&id)
	if err != nil {
		fmt.Printf("failed to insert chat || err: %v", err)
		return nil, err
	}

	return &desc.CreateChatResponse{Id: id}, nil
}

// DeleteChat удаление чата.
func (s *Server) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	pool := s.pgSQL.GetPool()

	_, err := pool.Exec(ctx, "DELETE FROM CHATS WHERE id = $1", req.Id)
	if err != nil {
		fmt.Printf("failed to delete user: %v", err)
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

// Start старт чат-сервера.
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.grpcConfig.Address())
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
