package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	pgt "github.com/jackc/pgx/v5/pgtype"
	db "shelke.dev/api/db/sqlc"
	"shelke.dev/api/internal/core/services"
)

type FeatureHandler struct {
	featureService *services.FeatureService
}

func NewFeatureHandler(featureService *services.FeatureService) *FeatureHandler {
	return &FeatureHandler{featureService: featureService}
}

// CreateFeature
// @Summary Create a new feature
// @Description Create a new feature with the provided details
// @Tags Features
// @Accept json
// @Produce json
// @Param feature body CreateFeatureRequest true "Feature creation request"
// @Success 201 {object} FeatureResponse
// @Failure 400 {string} string "Invalid request body or format"
// @Failure 500 {string} string "Failed to create feature"
// @Router /features [post]
func (h *FeatureHandler) CreateFeature(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateFeatureRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Printf("CreateFeature: Invalid request body: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("CreateFeature: Decoded request body: %+v\n", reqBody)

	arg := db.CreateFeatureParams{
		Name: reqBody.Name,
	}

	if reqBody.Description != nil {
		arg.Description = pgt.Text{String: *reqBody.Description, Valid: true}
	}
	if reqBody.CreatedBy != nil {
		createdByUUID, err := uuid.Parse(*reqBody.CreatedBy)
		if err != nil {
			fmt.Printf("CreateFeature: Invalid CreatedBy UUID: %v\n", err)
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

	feature, err := h.featureService.CreateFeature(r.Context(), arg)
	if err != nil {
		fmt.Printf("CreateFeature: Failed to create feature: %v\n", err)
		http.Error(w, "Failed to create feature", http.StatusInternalServerError)
		return
	}

	fmt.Printf("CreateFeature: Feature created successfully: %+v\n", feature)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(toFeatureResponse(feature))
}


// ListFeatures
// @Summary Get all features
// @Description Retrieve a list of all features
// @Tags Features
// @Produce json
// @Success 200 {array} FeatureResponse
// @Failure 500 {string} string "Failed to list features"
// @Router /features [get]
func (h *FeatureHandler) ListFeatures(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListFeatures: Calling service to list features")

	features, err := h.featureService.ListFeatures(r.Context())
	if err != nil {
		fmt.Printf("ListFeatures: Failed to list features: %v\n", err)
		http.Error(w, "Failed to list features", http.StatusInternalServerError)
		return
	}

	fmt.Printf("ListFeatures: Successfully listed %d features\n", len(features))

	featureResponses := make([]FeatureResponse, len(features))
	for i, feature := range features {
		featureResponses[i] = toFeatureResponse(feature)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(featureResponses)
}

// UpdateFeature
// @Summary Update an existing feature
// @Description Update an existing feature with the provided details
// @Tags Features
// @Accept json
// @Produce json
// @Param id path string true "Feature ID"
// @Param feature body UpdateFeatureRequest true "Feature update request"
// @Success 200 {object} FeatureResponse
// @Failure 400 {string} string "Invalid feature ID or request body"
// @Failure 500 {string} string "Failed to update feature"
// @Router /features/{id} [put]
func (h *FeatureHandler) UpdateFeature(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/features/")
	if idStr == "" {
		http.Error(w, "Feature ID is required", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Printf("UpdateFeature: Invalid feature ID: %v\n", err)
		http.Error(w, "Invalid feature ID", http.StatusBadRequest)
		return
	}

	var reqBody UpdateFeatureRequest

	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Printf("UpdateFeature: Invalid request body: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("UpdateFeature: Decoded request body: %+v\n", reqBody)

	arg := db.UpdateFeatureParams{
		ID: pgt.UUID{
			Bytes: id,
			Valid: true,
		},
		Name: reqBody.Name,
	}

	if reqBody.Description != nil {
		arg.Description = pgt.Text{String: *reqBody.Description, Valid: true}
	}
	if reqBody.Priority != nil {
		arg.Priority = pgt.Text{String: *reqBody.Priority, Valid: true}
	}
	if reqBody.Status != nil {
		arg.Status = pgt.Text{String: *reqBody.Status, Valid: true}
	}

	feature, err := h.featureService.UpdateFeature(r.Context(), arg)
	if err != nil {
		fmt.Printf("UpdateFeature: Failed to update feature: %v\n", err)
		http.Error(w, "Failed to update feature", http.StatusInternalServerError)
		return
	}

	fmt.Printf("UpdateFeature: Feature updated successfully: %+v\n", feature)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toFeatureResponse(feature))
}


// DeleteFeature
// @Summary Delete a feature
// @Description Delete a feature by its ID
// @Tags Features
// @Produce json
// @Param id path string true "Feature ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Invalid feature ID"
// @Failure 500 {string} string "Failed to delete feature"
// @Router /features/{id} [delete]
func (h *FeatureHandler) DeleteFeature(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/features/")
	if idStr == "" {
		http.Error(w, "Feature ID is required", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Printf("DeleteFeature: Invalid feature ID: %v\n", err)
		http.Error(w, "Invalid feature ID", http.StatusBadRequest)
		return
	}

	err = h.featureService.DeleteFeature(r.Context(), pgt.UUID{Bytes: id, Valid: true})
	if err != nil {
		fmt.Printf("DeleteFeature: Failed to delete feature: %v\n", err)
		http.Error(w, "Failed to delete feature", http.StatusInternalServerError)
		return
	}

	fmt.Printf("DeleteFeature: Feature deleted successfully: %s\n", id.String())

	w.WriteHeader(http.StatusNoContent)
}