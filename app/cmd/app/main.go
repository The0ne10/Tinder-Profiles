package main

import (
	"Profiles/app/internal/app"
	"Profiles/app/internal/config"
	"Profiles/app/internal/storage/postgres"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	LocalEnv = "local"
	DevEnv   = "dev"
	ProdEnv  = "prod"
)

func main() {
	// Подключениея конфига
	cfg := config.MustLoad()

	// Подключение логгирования
	logger := setupLogger(cfg.Env)

	// Запуск контекста
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Info("Показ конфига",
		"Config", cfg,
		"Storage", cfg.Storage,
		"Minio", cfg.Minio,
	)

	logger.Info("Логирование запущено...")

	if cfg.Env == LocalEnv {
		logger.Debug("Логгер дебаг запущен...")
	}

	// Подключения базы данных
	db, err := postgres.NewStorage(ctx, logger, cfg)
	if err != nil {
		logger.Error("Не удалось подключиться к базе данных", "err", err)
		os.Exit(1)
	}
	defer db.Close()

	// Создание и запуск приложения
	application := app.NewApp(ctx, cfg, logger, db)
	go func() {
		if err = application.Run(); err != nil {
			logger.Error("Ошибка при запуске приложения", slog.String("error", err.Error()))
			cancel()
		}
	}()

	// Ожидание сигнала для завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Остановка приложения
	application.Stop()
	logger.Info("Приложение остановлено")
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case LocalEnv:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case DevEnv:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case ProdEnv:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
