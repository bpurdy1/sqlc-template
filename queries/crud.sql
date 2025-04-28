-- name: InsertExample :one
INSERT INTO example_table (message)
VALUES (?)
RETURNING *;

-- name: SelectByIdExample :one
SELECT * FROM example_table WHERE id = ?;

-- name: ListExamples :many
SELECT * FROM example_table ORDER BY id ASC;

-- name: UpdateExample :one
UPDATE example_table
SET 
    message = ?, 
    modified_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: DeleteExample :one
DELETE FROM example_table
WHERE id = ?
RETURNING *;

-- name: UpsertExample :one
INSERT INTO example_table (id, message)
VALUES (?, ?)
ON CONFLICT (id) DO UPDATE SET
    message = EXCLUDED.message,
    modified_at = CURRENT_TIMESTAMP
RETURNING *;