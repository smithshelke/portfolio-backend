package httphandler

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	db "shelke.dev/api/db/sqlc"
)

func toFeatureResponse(feature db.Feature) FeatureResponse {
	response := FeatureResponse{
		ID:        uuid.UUID(feature.ID.Bytes).String(),
		Name:      feature.Name,
		CreatedAt: feature.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: feature.UpdatedAt.Time.Format(time.RFC3339),
	}
	if feature.Description.Valid {
		response.Description = &feature.Description.String
	}
	if feature.CreatedBy.Valid {
		createdBy := uuid.UUID(feature.CreatedBy.Bytes).String()
		response.CreatedBy = &createdBy
	}
	if feature.Priority.Valid {
		response.Priority = &feature.Priority.String
	}
	if feature.Status.Valid {
		response.Status = &feature.Status.String
	}
	return response
}

func toTaskResponse(task db.Task) TaskResponse {
	response := TaskResponse{
		ID:        uuid.UUID(task.ID.Bytes).String(),
		Name:      task.Name,
		CreatedAt: task.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: task.UpdatedAt.Time.Format(time.RFC3339),
		FeatureID: uuid.UUID(task.FeatureID.Bytes).String(),
	}
	if task.Description.Valid {
		response.Description = &task.Description.String
	}
	if task.CreatedBy.Valid {
		createdBy := uuid.UUID(task.CreatedBy.Bytes).String()
		response.CreatedBy = &createdBy
	}
	if task.FeatureName.Valid {
		response.FeatureName = &task.FeatureName.String
	}
	if task.Priority.Valid {
		response.Priority = &task.Priority.String
	}
	if task.Status.Valid {
		response.Status = &task.Status.String
	}
	if task.GitData != nil {
		response.GitData = json.RawMessage(task.GitData)
	}
	return response
}
