package dao

import (
	"database/sql"
	"fmt"
	"time"

	"assalielmehdi/eventify/db/model"
)

type ProcessDao struct {
	db *sql.DB
}

func NewProcessDao(db *sql.DB) *ProcessDao {
	return &ProcessDao{
		db: db,
	}
}

func (dao *ProcessDao) GetAll() ([]*model.Process, error) {
	q := `
		SELECT *
		FROM processes
		WHERE deleted_at = -1;
	`
	rows, err := query(dao.db, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ps := make([]*model.Process, 0)

	for rows.Next() {
		p, err := dao.scan(rows)
		if err != nil {
			return nil, err
		}

		ps = append(ps, p)
	}

	return ps, nil
}

func (dao *ProcessDao) GetByName(name string) (*model.Process, error) {
	q := `
		SELECT *
		FROM processes
		WHERE name = ? AND deleted_at = -1;
	`
	rows, err := query(dao.db, q, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, NotFoundError{fmt.Sprintf("Not Found: process{name=%s}", name)}
	}

	p, err := dao.scan(rows)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (dao *ProcessDao) Save(p *model.Process) error {
	_, err := dao.GetByName(p.Name)
	if err == nil { // exists
		return nil // nothing to update
	}

	return dao.insert(p)
}

func (dao *ProcessDao) DeleteByName(name string) error {
	_, err := dao.GetByName(name)
	if err != nil {
		return NotFoundError{fmt.Sprintf("Not Found: process{name=%s}", name)}
	}

	q := `
		UPDATE processes 
		SET deleted_at = ?
		WHERE name = ?;
	`
	_, err = exec(dao.db, q, time.Now().UnixMilli(), name)
	if err != nil {
		return err
	}

	return nil
}

func (dao *ProcessDao) insert(p *model.Process) error {
	q := `
		INSERT INTO processes(name, created_at, deleted_at)
		VALUES(?, ?, -1);
	`
	_, err := exec(dao.db, q, p.Name, time.Now().UnixMilli())
	if err != nil {
		return err
	}

	return nil
}

func (dao *ProcessDao) scan(rows *sql.Rows) (*model.Process, error) {
	var p model.Process
	err := rows.Scan(&p.Name, &p.CreatedAt, &p.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
