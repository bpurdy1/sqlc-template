package schema

import (
	"context"
	"database/sql"
	_ "embed"
)

//go:embed schema.sql
var Schema string

func CreateTableContext(ctx context.Context, db *sql.DB) (sql.Result, error) {
	return db.ExecContext(ctx, Schema)
}

func CreateTable(db *sql.DB) (sql.Result, error) {
	return db.Exec(Schema)
}
