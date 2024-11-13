package app

import (
	"context"

	"github.com/milovanovmaksim/chat-server/internal/closer"
	"github.com/milovanovmaksim/chat-server/internal/server/grpc"
)

type App struct {
	diContainer diContainer
	grpcServer  grpc.Server
}

// NeaApp создает новый объект App.
func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// Run запускает приложение.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.grpcServer.Start()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initGRPCServer,
		a.initdiContainer,
	}
	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(a.diContainer.GRPCConfig(), a.diContainer.ChatService(ctx))

	return nil
}

func (a *App) initdiContainer(_ context.Context) error {
	a.diContainer = newDiContainer()

	return nil
}
