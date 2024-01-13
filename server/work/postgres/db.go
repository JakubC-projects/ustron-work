package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDb(connectionString string) (*sql.DB, error) {
	return sql.Open("postgres", connectionString)
}
