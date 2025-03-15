package main

import (
	"Profiles/app/internal/app"
	"Profiles/app/internal/config"
	"Profiles/app/internal/storage/postgres"
	"context"
	"log/slog"
	"os"
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
	log := setupLogger(cfg.Env)

	// Запуск контекста
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Info("Показ конфига",
		"Config", cfg,
		"Storage", cfg.Storage,
		"Minio", cfg.Minio,
	)

	log.Info("Логирование запущено...")

	if cfg.Env == LocalEnv {
		log.Debug("Логгер дебаг запущен...")
	}

	// Подключения базы данных
	dbpool, err := postgres.NewStorage(ctx, log, cfg)
	if err != nil {
		log.Error("Не удалось подключиться к базе данных", "err", err)
		os.Exit(1)
	}

	// Запуск основного приложения
	application := app.NewApp(ctx, cfg, log, dbpool)
	if err := application.Run(); err != nil {
		log.Error("Ошибка запуска сервера %v", err)
	}
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
