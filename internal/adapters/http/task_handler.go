package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	pgt "github.com/jackc/pgx/v5/pgtype" // Import pgtype with an alias
	db "shelke.dev/api/db/sqlc"
	"shelke.dev/api/internal/core/services"
	// "github.com/sqlc-dev/pqtype" // No longer needed, pgtype.JSONB is used
	// "database/sql" // No longer needed, pgtype.Text is used
)

type TaskHandler struct {
	taskService    *services.TaskService
	featureService *services.FeatureService
}

func NewTaskHandler(taskService *services.TaskService, featureService *services.FeatureService) *TaskHandler {
	return &TaskHandler{taskService: taskService, featureService: featureService}
}

// CreateTask
// @Summary Create a new task
// @Description Create a new task with the provided details
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body CreateTaskRequest true "Task creation request"
// @Success 201 {object} TaskResponse
// @Failure 400 {string} string "Invalid request body or format"
// @Failure 500 {string} string "Failed to create task"
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Printf("CreateTask: Invalid request body: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	featureIDLog := "<nil>"
	if reqBody.FeatureID != nil {
		featureIDLog = *reqBody.FeatureID
	}
	createdByLog := "<nil>"
	if reqBody.CreatedBy != nil {
		createdByLog = *reqBody.CreatedBy
	}
	fmt.Printf("CreateTask: Decoded request body: Name:%s Description:%v CreatedBy:%s FeatureID:%s Priority:%v Status:%v GitData:%v\n",
		reqBody.Name, reqBody.Description, createdByLog, featureIDLog, reqBody.Priority, reqBody.Status, reqBody.GitData)

	arg := db.CreateTaskParams{
		Name: reqBody.Name,
	}

	var featureUUID uuid.UUID
	if reqBody.FeatureID != nil {
		featureUUID, err = uuid.Parse(*reqBody.FeatureID)
		if err != nil {
			fmt.Printf("CreateTask: Invalid FeatureID: %v\n", err)
			http.Error(w, "Invalid FeatureID format", http.StatusBadRequest)
			return
		}
		arg.FeatureID = pgt.UUID{Bytes: featureUUID, Valid: true}
	} else {
		fmt.Printf("CreateTask: FeatureID is required\n")
		http.Error(w, "FeatureID is required", http.StatusBadRequest)
		return
	}

	// Fetch feature name using featureService
	feature, err := h.featureService.GetFeature(r.Context(), pgt.UUID{Bytes: featureUUID, Valid: true})
	if err != nil {
		fmt.Printf("CreateTask: Failed to fetch feature for ID %s: %v\n", featureUUID.String(), err)
		http.Error(w, "Failed to fetch feature for provided FeatureID", http.StatusInternalServerError)
		return
	}
	arg.FeatureName = pgt.Text{String: feature.Name, Valid: true}

	if reqBody.Description != nil {
		arg.Description = pgt.Text{String: *reqBody.Description, Valid: true}
	}
	if reqBody.CreatedBy != nil {
		createdByUUID, err := uuid.Parse(*reqBody.CreatedBy)
		if err != nil {
			fmt.Printf("CreateTask: Invalid CreatedBy UUID: %v\n", err)
			http.Error(w, "Invalid CreatedBy UUID format", http.StatusBadRequest)
			return
		}
		arg.CreatedBy = pgt.UUID{Bytes: createdByUUID, Valid: true}
	}
	if reqBody.Priority != nil {
		arg.Priority = pgt.Text{String: *reqBody.Priority, Valid: true}
	}
	if reqBody.Status != nil {
		arg.Status = pgt.Text{String: *reqBody.Status, Valid: true}
	}
	if reqBody.GitData != nil { // Check if json.RawMessage is not nil
		arg.GitData = reqBody.GitData // Assign directly as []byte
	}

	fmt.Printf("CreateTask: Calling service with arguments: %+v\n", arg)

	task, err := h.taskService.CreateTask(r.Context(), arg)
	if err != nil {
		fmt.Printf("CreateTask: Failed to create task: %v\n", err)
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	fmt.Printf("CreateTask: Task created successfully: %+v\n", task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(toTaskResponse(task))
}

// ListTasks
// @Summary Get all tasks
// @Description Retrieve a list of all tasks
// @Tags Tasks
// @Produce json
// @Success 200 {array} TaskResponse
// @Failure 500 {string} string "Failed to list tasks"
// @Router /tasks [get]
func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListTasks: Calling service to list tasks")

	tasks, err := h.taskService.ListTasks(r.Context())
	if err != nil {
		fmt.Printf("ListTasks: Failed to list tasks: %v\n", err)
		http.Error(w, "Failed to list tasks", http.StatusInternalServerError)
		return
	}

	fmt.Printf("ListTasks: Successfully listed %d tasks\n", len(tasks))

	taskResponses := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		taskResponses[i] = toTaskResponse(task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskResponses)
}

// UpdateTask
// @Summary Update an existing task
// @Description Update an existing task with the provided details
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body UpdateTaskRequest true "Task update request"
// @Success 200 {object} TaskResponse
// @Failure 400 {string} string "Invalid task ID or request body"
// @Failure 500 {string} string "Failed to update task"
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if idStr == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Printf("UpdateTask: Invalid task ID: %v\n", err)
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var reqBody UpdateTaskRequest

	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Printf("UpdateTask: Invalid request body: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("UpdateTask: Decoded request body: %+v\n", reqBody)

	arg := db.UpdateTaskParams{
		ID: pgt.UUID{Bytes: id, Valid: true},
	}

	if reqBody.Name != nil {
		arg.Name = pgt.Text{String: *reqBody.Name, Valid: true}
	} else {
		arg.Name = pgt.Text{Valid: false}
	}
	if reqBody.Description != nil {
		arg.Description = pgt.Text{String: *reqBody.Description, Valid: true}
	} else {
		arg.Description = pgt.Text{Valid: false}
	}

	if reqBody.FeatureID != nil {
		featureUUID, err := uuid.Parse(*reqBody.FeatureID)
		if err != nil {
			fmt.Printf("UpdateTask: Invalid FeatureID: %v\n", err)
			http.Error(w, "Invalid FeatureID format", http.StatusBadRequest)
			return
		}
		arg.FeatureID = pgt.UUID{Bytes: featureUUID, Valid: true}

		// Fetch feature name using featureService
		feature, err := h.featureService.GetFeature(r.Context(), pgt.UUID{Bytes: featureUUID, Valid: true})
		if err != nil {
			fmt.Printf("UpdateTask: Failed to fetch feature for ID %s: %v\n", featureUUID.String(), err)
			http.Error(w, "Failed to fetch feature for provided FeatureID", http.StatusInternalServerError)
			return
		}
		arg.FeatureName = pgt.Text{String: feature.Name, Valid: true}
	} else {
		arg.FeatureID = pgt.UUID{Valid: false}
		arg.FeatureName = pgt.Text{Valid: false}
	}

	if reqBody.Priority != nil {
		arg.Priority = pgt.Text{String: *reqBody.Priority, Valid: true}
	} else {
		arg.Priority = pgt.Text{Valid: false}
	}
	if reqBody.Status != nil {
		arg.Status = pgt.Text{String: *reqBody.Status, Valid: true}
	} else {
		arg.Status = pgt.Text{Valid: false}
	}
	if reqBody.GitData != nil { // Check if json.RawMessage is not nil
		arg.GitData = reqBody.GitData // Assign directly as []byte
	} else {
		arg.GitData = nil // Explicitly set to nil if not provided
	}

	fmt.Printf("UpdateTask: Calling service with arguments: %+v\n", arg)

	task, err := h.taskService.UpdateTask(r.Context(), arg)
	if err != nil {
		fmt.Printf("UpdateTask: Failed to update task: %v\n", err)
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	fmt.Printf("UpdateTask: Task updated successfully: %+v\n", task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toTaskResponse(task))
}

// DeleteTask
// @Summary Delete a task
// @Description Delete a task by its ID
// @Tags Tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Invalid task ID"
// @Failure 500 {string} string "Failed to delete task"
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if idStr == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Printf("DeleteTask: Invalid task ID: %v\n", err)
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.taskService.DeleteTask(r.Context(), pgt.UUID{Bytes: id, Valid: true})
	if err != nil {
		fmt.Printf("DeleteTask: Failed to delete task: %v\n", err)
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	fmt.Printf("DeleteTask: Task deleted successfully: %s\n", id.String())

	w.WriteHeader(http.StatusNoContent)
}
