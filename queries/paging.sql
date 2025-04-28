-- name: ExamplesOffsertPage :many
SELECT *
FROM example_table
ORDER BY id ASC
LIMIT ?
OFFSET ?;

-- name: ExamplesCursorPage :many
SELECT *
FROM example_table
WHERE id > ?
ORDER BY id ASC
LIMIT ?;