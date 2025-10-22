-- name: CreateFeature :one
INSERT INTO features (
    id, name, description, created_at, updated_at, created_by, priority, status
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: ListFeatures :many
SELECT * FROM features
ORDER BY created_at DESC;

-- name: UpdateFeature :one
UPDATE features
SET
    name = $2,
    description = $3,
    updated_at = $4,
    created_by = $5,
    priority = $6,
    status = $7
WHERE id = $1
RETURNING *;

-- name: DeleteFeature :exec
DELETE FROM features
WHERE id = $1;

-- name: GetFeature :one
SELECT * FROM features
WHERE id = $1;