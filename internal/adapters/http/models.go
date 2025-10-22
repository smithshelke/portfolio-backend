package httphandler

import (
	"encoding/json"
)

// CreateTaskRequest represents the request body for creating a new task.
type CreateTaskRequest struct {
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	CreatedBy   *string         `json:"created_by"`
	FeatureID   *string         `json:"feature_id"`
	Priority    *string         `json:"priority"`
	Status      *string         `json:"status"`
	GitData     json.RawMessage `json:"git_data"`
}

// UpdateTaskRequest represents the request body for updating an existing task.
type UpdateTaskRequest struct {
	Name        *string         `json:"name"`
	Description *string         `json:"description"`
	FeatureID   *string         `json:"feature_id"`
	Priority    *string         `json:"priority"`
	Status      *string         `json:"status"`
	GitData     json.RawMessage `json:"git_data"`
}

// CreateFeatureRequest represents the request body for creating a new feature.
type CreateFeatureRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedBy   *string `json:"created_by"`
	Priority    *string `json:"priority"`
	Status      *string `json:"status"`
}

// UpdateFeatureRequest represents the request body for updating an existing feature.
type UpdateFeatureRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Priority    *string `json:"priority"`
	Status      *string `json:"status"`
}

// FeatureResponse represents the HTTP response for a feature.
type FeatureResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	CreatedBy   *string `json:"created_by,omitempty"`
	Priority    *string `json:"priority,omitempty"`
	Status      *string `json:"status,omitempty"`
}

// TaskResponse represents the HTTP response for a task.
type TaskResponse struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
	CreatedBy   *string         `json:"created_by,omitempty"`
	FeatureID   string          `json:"feature_id"`
	FeatureName *string         `json:"feature_name,omitempty"`
	Priority    *string         `json:"priority,omitempty"`
	Status      *string         `json:"status,omitempty"`
	GitData     json.RawMessage `json:"git_data,omitempty"`
}
