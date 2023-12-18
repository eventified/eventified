package test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"

	"assalielmehdi/eventify/db"
)

func setup(t *testing.T) (*assert.Assertions, *sql.DB, func(), error) {
	a := assert.New(t)

	f := fmt.Sprintf("%d_test.db", time.Now().UnixMicro())
	td := func() {
		os.Remove(f)
	}

	os.Setenv("DB_FILE", f)
	db, err := db.New()
	if err != nil {
		return a, nil, td, err
	}

	return a, db, td, nil
}
