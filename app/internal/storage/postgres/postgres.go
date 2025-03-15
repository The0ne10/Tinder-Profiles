package postgres

import (
	"Profiles/app/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"time"
)

// Подключение к базе данных
func NewStorage(
	ctx context.Context,
	logger *slog.Logger,
	cfg config.Config,
) (*pgxpool.Pool, error) {
	const op = "storage.NewStorage"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Storage.User, cfg.Storage.Password, cfg.Storage.Host, cfg.Storage.Port, cfg.Storage.Database)

	logger.Info(dsn)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Проверяем подключение
	if err = dbpool.Ping(ctx); err != nil {
		dbpool.Close() // Закрываем пул при ошибке
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("Подключение к базе данных установленно...")

	return dbpool, nil
}
