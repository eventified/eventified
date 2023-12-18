package migration

import "database/sql"

type Init struct{}

func (Init) Name() string {
	return "001_create_tables"
}

func (Init) Up(db *sql.Tx) error {
	query := `
		PRAGMA foreign_keys = ON;
		
		CREATE TABLE IF NOT EXISTS processes(
			name TEXT NOT NULL UNIQUE,
			created_at INTEGER NOT NULL,
			deleted_at INTEGER
		) STRICT;
		
		CREATE TABLE IF NOT EXISTS activities(
			name TEXT NOT NULL UNIQUE,
			process TEXT NOT NULL,
			created_at INTEGER NOT NULL,
			deleted_at INTEGER,
		
			FOREIGN KEY (process) REFERENCES processes(name)
			UNIQUE(name, process)
		) STRICT;
		
		CREATE TABLE IF NOT EXISTS connections(
			source TEXT NOT NULL,
			target TEXT NOT NULL,
			created_at INTEGER NOT NULL,
			deleted_at INTEGER,
		
			FOREIGN KEY (source) REFERENCES activities(name),
			FOREIGN KEY (target) REFERENCES activities(name),
			UNIQUE(source, target)
		) STRICT;	
	`

	_, err := db.Exec(query)

	return err
}

func (Init) Down(db *sql.Tx) error {
	return nil
}
