-- name: CreatePost :exec
INSERT INTO posts(title, url, description, published_at, feed_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
);

-- name: GetUserPosts :many
SELECT 
  p.*,
  f.name AS feed_name
FROM posts AS p
INNER JOIN feeds AS f
  ON p.feed_id = f.id
INNER JOIN feed_follows AS ff
  ON f.id = ff.feed_id
INNER JOIN users AS u
  ON ff.user_id = u.id
WHERE u.id = $1
ORDER BY p.published_at
LIMIT $2;
