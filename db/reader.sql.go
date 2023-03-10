// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: reader.sql

package db

import (
	"context"
)

const getData = `-- name: GetData :one
SELECT iri, jsonld
FROM data
WHERE iri=$1
`

func (q *Queries) GetData(ctx context.Context, iri string) (Datum, error) {
	row := q.db.QueryRowContext(ctx, getData, iri)
	var i Datum
	err := row.Scan(&i.Iri, &i.Jsonld)
	return i, err
}
