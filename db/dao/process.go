package dao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/eventified/eventified/db/model"
)

func GetProcessAll(db *sql.DB) ([]*model.Process, error) {
	q := `
		SELECT *
		FROM processes
		WHERE deleted_at = -1;
	`
	rows, err := query(db, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ps := make([]*model.Process, 0)

	for rows.Next() {
		p, err := scanProcess(rows)
		if err != nil {
			return nil, err
		}

		ps = append(ps, p)
	}

	return ps, nil
}

func GetProcessByName(db *sql.DB, name string) (*model.Process, error) {
	q := `
		SELECT *
		FROM processes
		WHERE name = ? AND deleted_at = -1;
	`
	rows, err := query(db, q, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, NotFoundError{fmt.Sprintf("Not Found: process{name=%s}", name)}
	}

	p, err := scanProcess(rows)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func SaveProcess(db *sql.DB, p *model.Process) error {
	_, err := GetProcessByName(db, p.Name)
	if err == nil { // exists
		return nil // nothing to update
	}

	return insertProcess(db, p)
}

func DeleteProcessByName(db *sql.DB, name string) error {
	_, err := GetProcessByName(db, name)
	if err != nil {
		return NotFoundError{fmt.Sprintf("Not Found: process{name=%s}", name)}
	}

	q := `
		UPDATE processes 
		SET deleted_at = ?
		WHERE name = ?;
	`
	_, err = exec(db, q, time.Now().UnixMilli(), name)
	if err != nil {
		return err
	}

	return nil
}

func insertProcess(db *sql.DB, p *model.Process) error {
	q := `
		INSERT INTO processes(name, created_at, deleted_at)
		VALUES(?, ?, -1);
	`
	_, err := exec(db, q, p.Name, time.Now().UnixMilli())
	if err != nil {
		return err
	}

	return nil
}

func scanProcess(rows *sql.Rows) (*model.Process, error) {
	var p model.Process
	err := rows.Scan(&p.Name, &p.CreatedAt, &p.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
