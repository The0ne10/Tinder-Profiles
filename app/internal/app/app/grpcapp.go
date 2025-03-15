package grpcapp

import (
	"Profiles/app/internal/config"
	"Profiles/app/internal/services"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

func StartGRPCServer(
	ctx context.Context,
	cfg config.Config,
	logger *slog.Logger,
) error {
	const op = "StartGRPCServer"

	lis, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		logger.Error("failed to listen", "address", cfg.Address)
	}

	grpcServer := grpc.NewServer()

	_, err = services.NewService(ctx, cfg, logger)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Регистрация сервисов (маршруты)
	//pb.RegisterMyServiceServer(grpcServer, servicesContainer.HelloService)

	logger.Info("gRPC server listening on " + cfg.Address)
	return grpcServer.Serve(lis)
}
