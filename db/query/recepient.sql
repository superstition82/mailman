-- name: GetRecepient :one
SELECT * FROM recepient
WHERE id = ? LIMIT 1;

-- name: ListRecepients :many
SELECT * FROM recepient
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CreateRecepient :one
INSERT INTO recepient (email, status) 
VALUES (?, ?)
RETURNING *;

-- name: DeleteRecepient :exec
DELETE FROM recepient
WHERE id = ?;