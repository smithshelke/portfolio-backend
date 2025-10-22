package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	db "shelke.dev/api/db/sqlc"
)

func NewDB() (*db.Queries, *pgxpool.Pool, error) {
	connStr := "postgres://user:password@localhost:5432/shelke_dev_api?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	queries := db.New(pool)

	return queries, pool, nil
}
