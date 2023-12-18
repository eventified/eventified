package dao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/eventified/eventified/db/model"
)

func GetActivityAll(db *sql.DB) ([]*model.Activity, error) {
	q := `
		SELECT *
		FROM activities
		WHERE deleted_at = -1;
	`
	rows, err := query(db, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	as := make([]*model.Activity, 0)

	for rows.Next() {
		a, err := scanActivity(rows)
		if err != nil {
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}

func GetActivityByName(db *sql.DB, name string) (*model.Activity, error) {
	q := `
		SELECT *
		FROM activities
		WHERE name = ? AND deleted_at = -1;
	`
	rows, err := query(db, q, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, NotFoundError{fmt.Sprintf("Not Found: activity{name=%s}", name)}
	}

	a, err := scanActivity(rows)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func SaveActivity(db *sql.DB, a *model.Activity) error {
	_, err := GetActivityByName(db, a.Name)
	if err == nil { // exists
		return nil // nothing to update
	}

	return insertActivity(db, a)
}

func DeleteActivityByName(db *sql.DB, name string) error {
	_, err := GetActivityByName(db, name)
	if err != nil {
		return NotFoundError{fmt.Sprintf("Not Found: activity{name=%s}", name)}
	}

	q := `
		UPDATE activities 
		SET deleted_at = ?
		WHERE name = ?;
	`
	_, err = exec(db, q, time.Now().UnixMilli(), name)
	if err != nil {
		return err
	}

	return nil
}

func insertActivity(db *sql.DB, a *model.Activity) error {
	q := `
		INSERT INTO activities(name, process, created_at, deleted_at)
		VALUES(?, ?, ?, -1);
	`
	_, err := exec(db, q, a.Name, a.Process, time.Now().UnixMilli())
	if err != nil {
		return err
	}

	return nil
}

func scanActivity(rows *sql.Rows) (*model.Activity, error) {
	var a model.Activity
	err := rows.Scan(&a.Name, &a.Process, &a.CreatedAt, &a.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
