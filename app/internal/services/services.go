package services

import (
	"Profiles/app/internal/config"
	minioService "Profiles/app/internal/services/minio"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type ServiceContainer struct {
	MinioService *minioService.MinioService
}

func New(
	ctx context.Context,
	cfg config.Config,
	logger *slog.Logger,
	db *pgxpool.Pool,
) (*ServiceContainer, error) {
	const op = "ServiceContainer.New"
	minio, err := minioService.New(ctx, cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &ServiceContainer{
		MinioService: minio,
	}, nil
}
