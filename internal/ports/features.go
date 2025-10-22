package ports

import (
	"context"

	pgt "github.com/jackc/pgx/v5/pgtype"
	db "shelke.dev/api/db/sqlc"
)

type FeatureService interface {
	CreateFeature(ctx context.Context, arg db.CreateFeatureParams) (db.Feature, error)
	ListFeatures(ctx context.Context) ([]db.Feature, error)
	UpdateFeature(ctx context.Context, arg db.UpdateFeatureParams) (db.Feature, error)
	DeleteFeature(ctx context.Context, id pgt.UUID) error
}
