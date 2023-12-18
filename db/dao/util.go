package dao

import (
	"database/sql"
	"fmt"

	"github.com/eventified/eventified/common"
)

func query(db *sql.DB, q string, args ...any) (*sql.Rows, *common.Error) {
	q = fmt.Sprintf("PRAGMA foreign_keys = ON;%s", q)

	rows, err := db.Query(q, args...)
	if err != nil {
		return nil, common.InternalError(err)
	}

	return rows, nil
}

func exec(db *sql.DB, q string, args ...any) (sql.Result, *common.Error) {
	q = fmt.Sprintf("PRAGMA foreign_keys = ON;%s", q)

	res, err := db.Exec(q, args...)
	if err != nil {
		return nil, common.InternalError(err)
	}

	return res, nil
}
