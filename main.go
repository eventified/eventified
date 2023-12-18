package main

import (
	"log/slog"
	"os"

	_ "modernc.org/sqlite"

	"assalielmehdi/eventify/db"
)

func main() {
	db, err := db.New()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

}
