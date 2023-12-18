package db

import (
	"context"
	"database/sql"
	"os"

	"github.com/eventified/eventified/db/migration"
)

func migrate(db *sql.DB) error {
	var migs = []migration.Migration{
		migration.Init{},
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, mig := range migs {
		err := mig.Up(tx)
		if err != nil {
			tx.Rollback()

			return err
		}
	}

	return tx.Commit()
}

func New() (*sql.DB, error) {
	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		dbFile = "db.sqlite3"
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = migrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
