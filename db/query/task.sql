-- name: CreateTask :one
INSERT INTO tasks (
    name,
    description,
    created_by,
    feature_id,
    feature_name,
    priority,
    status,
    git_data
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY created_at DESC;

-- name: UpdateTask :one
UPDATE tasks
SET
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description),
    feature_id = COALESCE(sqlc.narg(feature_id), feature_id),
    feature_name = COALESCE(sqlc.narg(feature_name), feature_name),
    priority = COALESCE(sqlc.narg(priority), priority),
    status = COALESCE(sqlc.narg(status), status),
    git_data = COALESCE(sqlc.narg(git_data), git_data)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;

-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1;