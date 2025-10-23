package db

import (
	"context"
	"fmt"
	"os" // Import the os package

	"github.com/jackc/pgx/v5/pgxpool"
	db "shelke.dev/api/db/sqlc"
)

func NewDB() (*db.Queries, *pgxpool.Pool, error) {
	connStr := os.Getenv("DB_URL") // Read from environment variable
	if connStr == "" {
		// Use default if DB_URL is not set
		connStr = "postgres://user:password@localhost:5432/shelke_dev_api?sslmode=disable"
	}

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	queries := db.New(pool)

	return queries, pool, nil
}
