package migration

import "database/sql"

type Migration interface {
	Name() string
	Up(*sql.Tx) error
	Down(*sql.Tx) error
}
