-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING id,
    created_at,
    updated_at,
    user_id,
    feed_id
)
SELECT iff.*,
  feeds.name as feed_name,
  users.name as user_name
FROM inserted_feed_follow as iff
  INNER JOIN feeds ON iff.feed_id = feeds.id
  INNER JOIN users ON iff.user_id = users.id;
-- name: GetFeedFollowsForUser :many
SELECT feeds.url,
  feeds.name as feed_name,
  users.name as user_name
FROM feed_follows as ff
  JOIN feeds ON ff.feed_id = feeds.id
  JOIN users ON ff.user_id = users.id
WHERE ff.user_id = $1;
-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows
WHERE user_id = $1
  AND feed_id = $2;