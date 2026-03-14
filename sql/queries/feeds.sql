-- name: CreateFeed :one
INSERT INTO feeds(name, url, user_id) 
VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetFeeds :many
SELECT 
  f.name as name,
  f.url as url,
  u.name as username
FROM feeds as f
INNER JOIN users as u
  ON f.user_id = u.id;

-- name: GetFeedID :one
SELECT id FROM feeds
WHERE url = $1;

-- name: CreateFeedFollow :one
WITH insert_feed_follow AS (
  INSERT INTO feed_follows(user_id, feed_id)
  VALUES (
    $1,
    $2
  )
  RETURNING *
)
SELECT
  i.*,
  u.name as user_name,
  f.name as feed_name
FROM insert_feed_follow as i
INNER JOIN users as u
  ON i.user_id = u.id
INNER JOIN feeds as f
  ON i.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT
  f.name as feed_name,
  u.name as user_name
FROM feed_follows as ff
INNER JOIN feeds as f
  ON ff.feed_id = f.id
INNER JOIN users as u
  ON ff.user_id = u.id
WHERE u.name = $1;

-- name: RemoveFeedFollow :exec
DELETE FROM feed_follows
WHERE feed_id = $1;
