-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetFeed :one
SELECT *
FROM feeds
WHERE feeds.url = $1;
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
-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = NOW(),
  last_fetched_at = NOW()
WHERE id = $1;
-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;