package app

import (
	grpcapp "Profiles/app/internal/app/app"
	"Profiles/app/internal/config"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

// Запуск основного приложения

type App struct {
	ctx    context.Context
	cfg    config.Config
	logger *slog.Logger
	db     *pgxpool.Pool
}

func NewApp(
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
	}
}

func (a *App) Run() error {
	const op = "App.Run"

	a.logger.Info("Запуск gRPC Сервера", slog.String("op", op))

	return grpcapp.StartGRPCServer(a.ctx, a.cfg, a.logger)
}
