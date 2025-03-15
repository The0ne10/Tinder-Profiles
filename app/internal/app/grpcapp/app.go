package grpcapp

import (
	"Profiles/app/internal/config"
	"Profiles/app/internal/services"
	helloService "Profiles/app/internal/services/hello_service"
	"context"
	"fmt"
	pb "github.com/The0ne10/myTinderProto/hello_service/proto"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	ctx    context.Context
	cfg    config.Config
	logger *slog.Logger
	db     *pgxpool.Pool
	server *grpc.Server
}

func New(
	ctx context.Context,
	cfg config.Config,
	logger *slog.Logger,
	db *pgxpool.Pool,
) *App {
	return &App{
		ctx:    ctx,
		cfg:    cfg,
		logger: logger,
		db:     db,
		server: grpc.NewServer(),
	}
}

func (a *App) Run() error {
	const op = "grpcapp.App.Run"

	lis, err := net.Listen("tcp", a.cfg.GRPC.Address)
	if err != nil {
		a.logger.Error("Ошибка при запуске gRPC сервера", slog.String("op", op), slog.String("error", err.Error()))
		return err
	}

	_, err = services.New(a.ctx, a.cfg, a.logger, a.db)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.logger.Info("gRPC сервер запущен", slog.String("address", a.cfg.GRPC.Address))

	pb.RegisterHelloServiceServer(a.server, helloService.New())

	return a.server.Serve(lis)
}

func (a *App) Stop() {
	const op = "grpcapp.App.Stop"

	a.logger.Info("Остановка gRPC сервера", slog.String("op", op))
	a.server.GracefulStop()
}
