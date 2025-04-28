package client

import (
	"database/sql"
	"sqlc-template/sqlc"
)

type Client interface {
	CustomSqlFunc()
}
type SqlClient struct {
	sqlc.Querier
	db *sql.DB
}

func NewSqlClient(db *sql.DB) Client {
	return &SqlClient{
		Querier: sqlc.New(db),
		db:      db,
	}
}
func (c *SqlClient) CustomSqlFunc() {
	// c.db
}
