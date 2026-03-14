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
JOIN users as u
  ON f.user_id = u.id;
