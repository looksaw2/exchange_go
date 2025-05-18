-- name: GetUserByID :one
SELECT * FROM  users 
WHERE id = sqlc.arg(user_id) LIMIT 1;


-- name: GetUserList :many
SELECT * FROM  users;


-- name: GetUserWithRange :many
SELECT * FROM  users
ORDER BY id
OFFSET sqlc.arg(page_offset)
LIMIT  sqlc.arg(page_limit);

-- name: GetUserByName :one
SELECT * FROM users
WHERE name = sqlc.arg(user_name);

-- name: CreateUser :one
INSERT INTO users (
    name,password
)VALUES(
    sqlc.arg(user_name) , sqlc.arg(user_password)
)RETURNING *;


-- name: UpdateUserById :one
UPDATE users
SET name = sqlc.arg(user_name) , password = sqlc.arg(user_password)
WHERE id = sqlc.arg(user_id)
RETURNING *;


-- name: UpdateUserByName :one
UPDATE users
SET password = sqlc.arg(uesr_password)
WHERE name = sqlc.arg(uesr_name)
RETURNING *;


-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = sqlc.arg(user_id);

-- name: DeleteUserByName :exec
DELETE FROM users
WHERE name = sqlc.arg(uesr_name);


-- name: GetExchangeRateByID :one
SELECT * FROM exchange_rate
WHERE id = sqlc.arg(id) LIMIT 1;


-- name: GetExchangeRateList :many
SELECT * FROM exchange_rate
ORDER BY id;


-- name: CreateExchangeRate :one
INSERT INTO exchange_rate(
    from_currency , to_currency,rate
)VALUES(
    $1,$2,$3
)RETURNING *;


-- name: GetArticleByID :one
SELECT * FROM article 
WHERE id = $1 LIMIT 1;

-- name: GetArticleList :many
SELECT * FROM article
ORDER BY id;

-- name: CreateArticle :one
INSERT INTO article (
    title,content,preview
)VALUES(
    $1,$2,$3
)RETURNING *;


-- name: UpdateArticle :one
UPDATE article
SET title = $2 , content = $3 , preview = $4
WHERE id = $1
RETURNING *;


-- name: DeleteARticle :exec
DELETE FROM article 
WHERE id = $1;

