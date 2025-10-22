package ports

import (
	"context"

	pgt "github.com/jackc/pgx/v5/pgtype"
	db "shelke.dev/api/db/sqlc"
)

type TaskService interface {
	CreateTask(ctx context.Context, arg db.CreateTaskParams) (db.Task, error)
	ListTasks(ctx context.Context) ([]db.Task, error)
	UpdateTask(ctx context.Context, arg db.UpdateTaskParams) (db.Task, error)
	DeleteTask(ctx context.Context, id pgt.UUID) error
}
