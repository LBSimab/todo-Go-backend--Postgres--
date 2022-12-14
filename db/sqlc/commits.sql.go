// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: commits.sql

package db

import (
	"context"
	"database/sql"
)

const createcommit = `-- name: Createcommit :one
INSERT INTO commits(
  title, supervisor_id,task_id,comment,category,user_id
) VALUES (
  $1, $2, $3, $4 , $5, $6
)
RETURNING title, commit_id, user_id, task_id, comment, category, supervisor_id
`

type CreatecommitParams struct {
	Title        sql.NullString `json:"title"`
	SupervisorID sql.NullInt32  `json:"supervisor_id"`
	TaskID       sql.NullInt32  `json:"task_id"`
	Comment      sql.NullString `json:"comment"`
	Category     sql.NullString `json:"category"`
	UserID       sql.NullInt32  `json:"user_id"`
}

func (q *Queries) Createcommit(ctx context.Context, arg CreatecommitParams) (Commit, error) {
	row := q.db.QueryRowContext(ctx, createcommit,
		arg.Title,
		arg.SupervisorID,
		arg.TaskID,
		arg.Comment,
		arg.Category,
		arg.UserID,
	)
	var i Commit
	err := row.Scan(
		&i.Title,
		&i.CommitID,
		&i.UserID,
		&i.TaskID,
		&i.Comment,
		&i.Category,
		&i.SupervisorID,
	)
	return i, err
}

const deletecommit = `-- name: Deletecommit :exec
DELETE FROM commits
WHERE commit_id = $1
`

func (q *Queries) Deletecommit(ctx context.Context, commitID int32) error {
	_, err := q.db.ExecContext(ctx, deletecommit, commitID)
	return err
}

const getcommit = `-- name: Getcommit :one
SELECT title, commit_id, user_id, task_id, comment, category, supervisor_id FROM commits
WHERE commit_id = $1 LIMIT 1
`

func (q *Queries) Getcommit(ctx context.Context, commitID int32) (Commit, error) {
	row := q.db.QueryRowContext(ctx, getcommit, commitID)
	var i Commit
	err := row.Scan(
		&i.Title,
		&i.CommitID,
		&i.UserID,
		&i.TaskID,
		&i.Comment,
		&i.Category,
		&i.SupervisorID,
	)
	return i, err
}

const listcommits = `-- name: Listcommits :many
SELECT title, commit_id, user_id, task_id, comment, category, supervisor_id FROM commits
ORDER BY title
`

func (q *Queries) Listcommits(ctx context.Context) ([]Commit, error) {
	rows, err := q.db.QueryContext(ctx, listcommits)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Commit
	for rows.Next() {
		var i Commit
		if err := rows.Scan(
			&i.Title,
			&i.CommitID,
			&i.UserID,
			&i.TaskID,
			&i.Comment,
			&i.Category,
			&i.SupervisorID,
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

const updatecommit = `-- name: Updatecommit :exec
UPDATE commits
set title= $2,
comment = $3
WHERE commit_id = $1
`

type UpdatecommitParams struct {
	CommitID int32          `json:"commit_id"`
	Title    sql.NullString `json:"title"`
	Comment  sql.NullString `json:"comment"`
}

func (q *Queries) Updatecommit(ctx context.Context, arg UpdatecommitParams) error {
	_, err := q.db.ExecContext(ctx, updatecommit, arg.CommitID, arg.Title, arg.Comment)
	return err
}
