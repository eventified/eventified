package dao

import (
	"database/sql"
	"fmt"
)

func query(db *sql.DB, q string, args ...any) (*sql.Rows, error) {
	q = fmt.Sprintf("PRAGMA foreign_keys = ON;%s", q)
	return db.Query(q, args...)
}

func exec(db *sql.DB, q string, args ...any) (sql.Result, error) {
	q = fmt.Sprintf("PRAGMA foreign_keys = ON;%s", q)
	return db.Exec(q, args...)
}
