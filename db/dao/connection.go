package dao

// import (
// 	"database/sql"
// 	"time"

// 	"assalielmehdi/eventify/db/model"
// )

// type ConnectionDao struct {
// 	defaultDao[model.Connection]
// }

// func scanConnection(rows *sql.Rows, rec *model.Connection) error {
// 	var createdAt, deletedAt int64

// 	err := rows.Scan(
// 		&rec.Id,
// 		&rec.FromActivityId,
// 		&rec.ToActivityId,
// 		&createdAt,
// 		&deletedAt,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	rec.CreatedAt = time.UnixMilli(createdAt)
// 	rec.DeletedAt = time.UnixMilli(deletedAt)

// 	return nil
// }

// func insertConnection(db *sql.DB, rec *model.Connection) error {
// 	return nil
// }

// func updateConnection(db *sql.DB, rec *model.Connection) error {
// 	return nil
// }

// func NewConnectionDao(db *sql.DB) *ConnectionDao {
// 	return &ConnectionDao{
// 		defaultDao: defaultDao[model.Connection]{
// 			db:     db,
// 			table:  "connections",
// 			scan:   scanConnection,
// 			insert: insertConnection,
// 			update: updateConnection,
// 		},
// 	}
// }
