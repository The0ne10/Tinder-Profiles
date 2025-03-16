package profileService

import (
	"Profiles/app/internal/models"
	"Profiles/app/internal/repository/profile_repository"
	"context"
	"fmt"
	pb "github.com/The0ne10/myTinderProto/profile_service/proto"
	"log/slog"
)

type Server struct {
	pb.UnimplementedProfileServiceServer
	ctx               context.Context
	logger            *slog.Logger
	ProfileRepository *profileRepository.ProfileRepository
}

func New(
	ctx context.Context,
	logger *slog.Logger,
	ProfileRepository *profileRepository.ProfileRepository,
) *Server {
	return &Server{
		ctx:               ctx,
		logger:            logger,
		ProfileRepository: ProfileRepository,
	}
}

func (s *Server) CreateProfile(
	ctx context.Context,
	req *pb.CreateProfileRequest,
) (*pb.CreateProfileResponse, error) {
	const op = "gRPC.CreateProfile"

	// TODO: Обновить протофайл чтобы можно было загружать еще и фото
	profile := &models.Profile{
		Name:      req.GetName(),
		Latitude:  req.GetLatitude(),
		Longitude: req.GetLongitude(),
	}

	userId, err := s.ProfileRepository.CreateProfile(ctx, profile)
	if err != nil {
		s.logger.Error("Ошибка создания профиля", "err", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// TODO: добавить сервис который бы загружал фотографии профиля в минио

	s.logger.Info("Профиль успешно создан...", "user_id", userId, "Name", req.GetName())

	return &pb.CreateProfileResponse{
		UserId:  userId,
		Message: "Profile successfully created",
	}, nil
}

func (s *Server) GetProfile(
	ctx context.Context,
	req *pb.GetProfileRequest,
) (*pb.GetProfileResponse, error) {
	panic("implement me")
}

func (s *Server) UpdateProfile(
	ctx context.Context,
	req *pb.UpdateProfileRequest,
) (*pb.UpdateProfileResponse, error) {
	panic("implement me")
}

func (s *Server) DeleteProfile(
	ctx context.Context,
	req *pb.DeleteProfileRequest,
) (*pb.DeleteProfileResponse, error) {
	panic("implement me")
}
