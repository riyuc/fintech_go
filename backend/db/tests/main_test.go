package db_test

import (
	"database/sql"
	db "github/riyuc/fintech_backend/db/sqlc"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

var testQuery *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open("postgres","postgres://root:secret@localhost:5432/fintech_db?sslmode=disable")
	if err != nil {
		log.Fatal("Could not connect to the database", err)
	}

	testQuery = db.New(conn)

	os.Exit(m.Run())
}