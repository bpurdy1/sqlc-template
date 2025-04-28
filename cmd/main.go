package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sqlc-template/schema"
	"sqlc-template/sqlc"

	_ "modernc.org/sqlite"
)

var ctx = context.Background()

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	// db, err := sql.Open("sqlite", "data.db?_busy_timeout=1500_journal=WAL&_timeout=5000&cache=shared")
	// db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?_busy_timeout=1500_journal=WAL&_timeout=5000&cache=shared", dbName))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = schema.CreateTable(db)
	if err != nil {
		panic(err)
	}
	querier, err := sqlc.Prepare(ctx, db)
	if err != nil {
		panic(err)
	}
	_, err = querier.InsertExample(ctx, "Hello, World!")
	if err != nil {
		panic(err)
	}
	rows, err := querier.ListExamples(ctx)
	for _, r := range rows {
		b, _ := json.MarshalIndent(r, "", "  ")
		fmt.Println(string(b))
	}
}
