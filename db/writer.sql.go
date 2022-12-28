// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: writer.sql

package db

import (
	"context"
)

const postData = `-- name: PostData :exec
INSERT INTO data (iri, context, jsonld)
VALUES ($1, $2, $3)
`

type PostDataParams struct {
	Iri     string
	Context string
	Jsonld  string
}

func (q *Queries) PostData(ctx context.Context, arg PostDataParams) error {
	_, err := q.db.ExecContext(ctx, postData, arg.Iri, arg.Context, arg.Jsonld)
	return err
}
