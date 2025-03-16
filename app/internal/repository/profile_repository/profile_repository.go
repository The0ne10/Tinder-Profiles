package profileRepository

import (
	"Profiles/app/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type ProfileRepository struct {
	db     *pgxpool.Pool
	ctx    context.Context
	logger *slog.Logger
}

func New(
	ctx context.Context,
	logger *slog.Logger,
	db *pgxpool.Pool,
) *ProfileRepository {
	return &ProfileRepository{
		db:     db,
		logger: logger,
		ctx:    ctx,
	}
}

func (p *ProfileRepository) CreateProfile(
	ctx context.Context,
	profile *models.Profile,
) (int64, error) {
	const op = "ProfileRepository.CreateProfile"

	query := `INSERT INTO profiles (name, latitude, longitude) VALUES ($1, $2, $3) RETURNING id`

	err := p.db.QueryRow(ctx, query, profile.Name, profile.Latitude, profile.Longitude).Scan(&profile.ID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return profile.ID, nil
}
