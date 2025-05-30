// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package sqlc_gen

import (
	"context"
	"database/sql"
)

const createArticle = `-- name: CreateArticle :one
INSERT INTO article (
    title,content,preview
)VALUES(
    $1,$2,$3
)RETURNING id, title, content, preview
`

type CreateArticleParams struct {
	Title   sql.NullString `json:"title"`
	Content sql.NullString `json:"content"`
	Preview sql.NullString `json:"preview"`
}

func (q *Queries) CreateArticle(ctx context.Context, arg CreateArticleParams) (Article, error) {
	row := q.queryRow(ctx, q.createArticleStmt, createArticle, arg.Title, arg.Content, arg.Preview)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Preview,
	)
	return i, err
}

const createExchangeRate = `-- name: CreateExchangeRate :one
INSERT INTO exchange_rate(
    from_currency , to_currency,rate
)VALUES(
    $1,$2,$3
)RETURNING id, from_currency, to_currency, rate, date
`

type CreateExchangeRateParams struct {
	FromCurrency sql.NullString  `json:"from_currency"`
	ToCurrency   sql.NullString  `json:"to_currency"`
	Rate         sql.NullFloat64 `json:"rate"`
}

func (q *Queries) CreateExchangeRate(ctx context.Context, arg CreateExchangeRateParams) (ExchangeRate, error) {
	row := q.queryRow(ctx, q.createExchangeRateStmt, createExchangeRate, arg.FromCurrency, arg.ToCurrency, arg.Rate)
	var i ExchangeRate
	err := row.Scan(
		&i.ID,
		&i.FromCurrency,
		&i.ToCurrency,
		&i.Rate,
		&i.Date,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    name,password
)VALUES(
    $1 , $2
)RETURNING id, name, password
`

type CreateUserParams struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.UserName, arg.UserPassword)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Password)
	return i, err
}

const deleteARticle = `-- name: DeleteARticle :exec
DELETE FROM article 
WHERE id = $1
`

func (q *Queries) DeleteARticle(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteARticleStmt, deleteARticle, id)
	return err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, userID int64) error {
	_, err := q.exec(ctx, q.deleteUserByIDStmt, deleteUserByID, userID)
	return err
}

const deleteUserByName = `-- name: DeleteUserByName :exec
DELETE FROM users
WHERE name = $1
`

func (q *Queries) DeleteUserByName(ctx context.Context, uesrName string) error {
	_, err := q.exec(ctx, q.deleteUserByNameStmt, deleteUserByName, uesrName)
	return err
}

const getArticleByID = `-- name: GetArticleByID :one
SELECT id, title, content, preview FROM article 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetArticleByID(ctx context.Context, id int64) (Article, error) {
	row := q.queryRow(ctx, q.getArticleByIDStmt, getArticleByID, id)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Preview,
	)
	return i, err
}

const getArticleList = `-- name: GetArticleList :many
SELECT id, title, content, preview FROM article
ORDER BY id
`

func (q *Queries) GetArticleList(ctx context.Context) ([]Article, error) {
	rows, err := q.query(ctx, q.getArticleListStmt, getArticleList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Article
	for rows.Next() {
		var i Article
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.Preview,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExchangeRateByID = `-- name: GetExchangeRateByID :one
SELECT id, from_currency, to_currency, rate, date FROM exchange_rate
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetExchangeRateByID(ctx context.Context, id int64) (ExchangeRate, error) {
	row := q.queryRow(ctx, q.getExchangeRateByIDStmt, getExchangeRateByID, id)
	var i ExchangeRate
	err := row.Scan(
		&i.ID,
		&i.FromCurrency,
		&i.ToCurrency,
		&i.Rate,
		&i.Date,
	)
	return i, err
}

const getExchangeRateList = `-- name: GetExchangeRateList :many
SELECT id, from_currency, to_currency, rate, date FROM exchange_rate
ORDER BY id
`

func (q *Queries) GetExchangeRateList(ctx context.Context) ([]ExchangeRate, error) {
	rows, err := q.query(ctx, q.getExchangeRateListStmt, getExchangeRateList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ExchangeRate
	for rows.Next() {
		var i ExchangeRate
		if err := rows.Scan(
			&i.ID,
			&i.FromCurrency,
			&i.ToCurrency,
			&i.Rate,
			&i.Date,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, name, password FROM  users 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, userID int64) (User, error) {
	row := q.queryRow(ctx, q.getUserByIDStmt, getUserByID, userID)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Password)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT id, name, password FROM users
WHERE name = $1
`

func (q *Queries) GetUserByName(ctx context.Context, userName string) (User, error) {
	row := q.queryRow(ctx, q.getUserByNameStmt, getUserByName, userName)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Password)
	return i, err
}

const getUserList = `-- name: GetUserList :many
SELECT id, name, password FROM  users
`

func (q *Queries) GetUserList(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.getUserListStmt, getUserList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.Password); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserWithRange = `-- name: GetUserWithRange :many
SELECT id, name, password FROM  users
ORDER BY id
OFFSET $1
LIMIT  $2
`

type GetUserWithRangeParams struct {
	PageOffset int32 `json:"page_offset"`
	PageLimit  int32 `json:"page_limit"`
}

func (q *Queries) GetUserWithRange(ctx context.Context, arg GetUserWithRangeParams) ([]User, error) {
	rows, err := q.query(ctx, q.getUserWithRangeStmt, getUserWithRange, arg.PageOffset, arg.PageLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Name, &i.Password); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateArticle = `-- name: UpdateArticle :one
UPDATE article
SET title = $2 , content = $3 , preview = $4
WHERE id = $1
RETURNING id, title, content, preview
`

type UpdateArticleParams struct {
	ID      int64          `json:"id"`
	Title   sql.NullString `json:"title"`
	Content sql.NullString `json:"content"`
	Preview sql.NullString `json:"preview"`
}

func (q *Queries) UpdateArticle(ctx context.Context, arg UpdateArticleParams) (Article, error) {
	row := q.queryRow(ctx, q.updateArticleStmt, updateArticle,
		arg.ID,
		arg.Title,
		arg.Content,
		arg.Preview,
	)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Preview,
	)
	return i, err
}

const updateUserById = `-- name: UpdateUserById :one
UPDATE users
SET name = $1 , password = $2
WHERE id = $3
RETURNING id, name, password
`

type UpdateUserByIdParams struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserID       int64  `json:"user_id"`
}

func (q *Queries) UpdateUserById(ctx context.Context, arg UpdateUserByIdParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserByIdStmt, updateUserById, arg.UserName, arg.UserPassword, arg.UserID)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Password)
	return i, err
}

const updateUserByName = `-- name: UpdateUserByName :one
UPDATE users
SET password = $1
WHERE name = $2
RETURNING id, name, password
`

type UpdateUserByNameParams struct {
	UesrPassword string `json:"uesr_password"`
	UesrName     string `json:"uesr_name"`
}

func (q *Queries) UpdateUserByName(ctx context.Context, arg UpdateUserByNameParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserByNameStmt, updateUserByName, arg.UesrPassword, arg.UesrName)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Password)
	return i, err
}
