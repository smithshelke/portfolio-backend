package services

import (
	"context"
	"fmt"

	pgt "github.com/jackc/pgx/v5/pgtype"
	db "shelke.dev/api/db/sqlc"
)

type TaskService struct {
	queries *db.Queries
}

func NewTaskService(queries *db.Queries) *TaskService {
	return &TaskService{queries: queries}
}

func (s *TaskService) CreateTask(ctx context.Context, arg db.CreateTaskParams) (db.Task, error) {
	fmt.Printf("TaskService: Creating task with arguments: %+v\n", arg)
	task, err := s.queries.CreateTask(ctx, arg)
	if err != nil {
		fmt.Printf("TaskService: Failed to create task: %v\n", err)
		return db.Task{}, err
	}
	fmt.Printf("TaskService: Task created successfully: %+v\n", task)
	return task, nil
}

func (s *TaskService) ListTasks(ctx context.Context) ([]db.Task, error) {
	fmt.Println("TaskService: Listing tasks")
	tasks, err := s.queries.ListTasks(ctx)
	if err != nil {
		fmt.Printf("TaskService: Failed to list tasks: %v\n", err)
		return nil, err
	}
	fmt.Printf("TaskService: Successfully listed %d tasks\n", len(tasks))
	return tasks, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, arg db.UpdateTaskParams) (db.Task, error) {
	fmt.Printf("TaskService: Updating task with arguments: %+v\n", arg)
	task, err := s.queries.UpdateTask(ctx, arg)
	if err != nil {
		fmt.Printf("TaskService: Failed to update task: %v\n", err)
		return db.Task{}, err
	}
	fmt.Printf("TaskService: Task updated successfully: %+v\n", task)
	return task, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id pgt.UUID) error {
	fmt.Printf("TaskService: Deleting task with ID: %v\n", id)
	err := s.queries.DeleteTask(ctx, id)
	if err != nil {
		fmt.Printf("TaskService: Failed to delete task: %v\n", err)
		return fmt.Errorf("failed to delete task: %w", err)
	}
	fmt.Printf("TaskService: Task deleted successfully: %v\n", id)
	return nil
}
