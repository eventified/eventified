package main

import (
	"log/slog"
	"os"

	_ "modernc.org/sqlite"

	"github.com/eventified/eventified/db"
)

func main() {
	db, err := db.New()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

}
