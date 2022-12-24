// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: reader.sql

package db

import (
	"context"
)

const getData = `-- name: GetData :one
SELECT id, canon, context, jsonld
FROM data
WHERE id=$1
`

func (q *Queries) GetData(ctx context.Context, id int32) (Datum, error) {
	row := q.db.QueryRowContext(ctx, getData, id)
	var i Datum
	err := row.Scan(
		&i.ID,
		&i.Canon,
		&i.Context,
		&i.Jsonld,
	)
	return i, err
}
