package minioService

import (
	"Profiles/app/internal/config"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log/slog"
)

type MinioService struct {
	client *minio.Client
	bucket string
}

func New(
	ctx context.Context,
	cfg config.Config,
	logger *slog.Logger,
) (*MinioService, error) {
	const op = "minio.New"

	endPoint := fmt.Sprintf("%s:%s", cfg.Minio.Host, cfg.Minio.Port)

	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.User, cfg.Minio.Password, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	exists, err := minioClient.BucketExists(ctx, cfg.Minio.Bucket)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if !exists {
		err = minioClient.MakeBucket(ctx, cfg.Minio.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("не удалось создать бакет: %v", err)
		}
	}

	logger.Info("Подключение к Minio успешно установленно...",
		"end-point", endPoint,
		"bucket", cfg.Minio.Bucket,
	)

	return &MinioService{
		client: minioClient,
		bucket: cfg.Minio.Bucket,
	}, nil
}
