package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	pgt "github.com/jackc/pgx/v5/pgtype"
	db "shelke.dev/api/db/sqlc"
)

type FeatureService struct {
	queries *db.Queries
}

func NewFeatureService(queries *db.Queries) *FeatureService {
	return &FeatureService{queries: queries}
}

func (s *FeatureService) CreateFeature(ctx context.Context, arg db.CreateFeatureParams) (db.Feature, error) {
	// Generate a new UUID for the feature
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return db.Feature{}, fmt.Errorf("failed to generate UUID: %w", err)
	}
	arg.ID = pgt.UUID{Bytes: newUUID, Valid: true}
	arg.CreatedAt = pgt.Timestamptz{Time: time.Now(), Valid: true}
	arg.UpdatedAt = pgt.Timestamptz{Time: time.Now(), Valid: true}

	feature, err := s.queries.CreateFeature(ctx, arg)
	if err != nil {
		return db.Feature{}, fmt.Errorf("failed to create feature: %w", err)
	}
	return feature, nil
}

func (s *FeatureService) ListFeatures(ctx context.Context) ([]db.Feature, error) {
	features, err := s.queries.ListFeatures(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list features: %w", err)
	}
	return features, nil
}

func (s *FeatureService) UpdateFeature(ctx context.Context, arg db.UpdateFeatureParams) (db.Feature, error) {
	arg.UpdatedAt = pgt.Timestamptz{Time: time.Now(), Valid: true}
	feature, err := s.queries.UpdateFeature(ctx, arg)
	if err != nil {
		return db.Feature{}, fmt.Errorf("failed to update feature: %w", err)
	}
	return feature, nil
}

func (s *FeatureService) DeleteFeature(ctx context.Context, id pgt.UUID) error {
	err := s.queries.DeleteFeature(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete feature: %w", err)
	}
	return nil
}

func (s *FeatureService) GetFeature(ctx context.Context, id pgt.UUID) (db.Feature, error) {
	feature, err := s.queries.GetFeature(ctx, id)
	if err != nil {
		return db.Feature{}, fmt.Errorf("failed to get feature: %w", err)
	}
	return feature, nil
}
