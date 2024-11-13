package app

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/milovanovmaksim/chat-server/internal/closer"
	"github.com/milovanovmaksim/chat-server/internal/server/grpc"
)

// App чат-приложение.
type App struct {
	diContainer diContainer
	grpcServer  grpc.Server
	envPath     string
}

// NeaApp создает новый объект App.
func NewApp(ctx context.Context, envPath string) (*App, error) {
	app := &App{envPath: envPath}

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

func (a *App) initConfig(_ctx context.Context) error {
	err := godotenv.Load(a.envPath)
	if err != nil {
		log.Printf("failed to load config || err: %v", err)
		return err
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initGRPCServer,
		a.initDiContainer,
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

func (a *App) initDiContainer(_ context.Context) error {
	a.diContainer = newDiContainer()

	return nil
}
