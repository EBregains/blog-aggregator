-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetFeeds :many
SELECT *
FROM feeds;
-- name: GetFeedsWithUsernames :many
SELECT f.id,
  f.created_at,
  f.updated_at,
  f.name,
  f.url,
  u.name as user_name
FROM feeds as f
  LEFT JOIN users as u ON f.user_id = u.id;