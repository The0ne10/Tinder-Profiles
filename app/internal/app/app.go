package app

import (
	"Profiles/app/internal/app/grpcapp"
	"Profiles/app/internal/config"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

// Запуск основного приложения

type App struct {
	ctx        context.Context
	cfg        config.Config
	logger     *slog.Logger
	db         *pgxpool.Pool
	grpcServer *grpcapp.App
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

	// Инициализация gRPC сервера
	a.grpcServer = grpcapp.New(a.ctx, a.cfg, a.logger, a.db)

	// Инициализация HTTP сервера Пример
	//a.httpServer = httpapp.NewHTTPApp(a.ctx, a.cfg, a.logger, a.db)
	// Запуск gRPC сервера в отдельной горутине

	go func() {
		a.logger.Info("Запуск gRPC сервера", slog.String("op", op))
		if err := a.grpcServer.Run(); err != nil {
			a.logger.Error("Ошибка при запуске gRPC сервера", slog.String("op", op), slog.String("error", err.Error()))
		}
	}()

	return nil
}

func (a *App) Stop() {
	const op = "App.Stop"

	// Остановка gRPC сервера
	if a.grpcServer != nil {
		a.logger.Info("Остановка gRPC сервера", slog.String("op", op))
		a.grpcServer.Stop()
	}

	// Остановка HTTP сервера
	//if a.httpServer != nil {
	//	a.logger.Info("Остановка HTTP сервера", slog.String("op", op))
	//	if err := a.httpServer.Stop(); err != nil {
	//		a.logger.Error("Ошибка при остановке HTTP сервера", slog.String("op", op), slog.String("error", err.Error()))
	//	}
	//}
}
