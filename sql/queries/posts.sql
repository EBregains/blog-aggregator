-- name: CreatePost :one
INSERT INTO posts (
    id,
    created_at,
    updated_at,
    title,
    url,
    description,
    published_at,
    feed_id
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;
-- name: GetPostsForUser :many
WITH all_feeds AS (
  SELECT *
  FROM feed_follows
  WHERE user_id = $1
)
SELECT (
    posts.id,
    posts.title,
    posts.description,
    posts.url,
    posts.published_at,
    posts.feed_id
  )
FROM all_feeds
  INNER JOIN posts ON posts.feed_id = all_feeds.feed_id
ORDER BY posts.published_at DESC
LIMIT $2;