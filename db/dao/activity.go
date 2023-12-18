package dao

import (
	"database/sql"
	"fmt"
	"time"

	"assalielmehdi/eventify/db/model"
)

type ActivityDao struct {
	db *sql.DB
}

func NewActivityDao(db *sql.DB) *ActivityDao {
	return &ActivityDao{
		db: db,
	}
}

func (dao *ActivityDao) GetAll() ([]*model.Activity, error) {
	q := `
		SELECT *
		FROM activities
		WHERE deleted_at = -1;
	`
	rows, err := query(dao.db, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	as := make([]*model.Activity, 0)

	for rows.Next() {
		a, err := dao.scan(rows)
		if err != nil {
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}

func (dao *ActivityDao) GetByName(name string) (*model.Activity, error) {
	q := `
		SELECT *
		FROM activities
		WHERE name = ? AND deleted_at = -1;
	`
	rows, err := query(dao.db, q, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, NotFoundError{fmt.Sprintf("Not Found: activity{name=%s}", name)}
	}

	a, err := dao.scan(rows)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (dao *ActivityDao) Save(a *model.Activity) error {
	_, err := dao.GetByName(a.Name)
	if err == nil { // exists
		return nil // nothing to update
	}

	return dao.insert(a)
}

func (dao *ActivityDao) DeleteByName(name string) error {
	_, err := dao.GetByName(name)
	if err != nil {
		return NotFoundError{fmt.Sprintf("Not Found: activity{name=%s}", name)}
	}

	q := `
		UPDATE activities 
		SET deleted_at = ?
		WHERE name = ?;
	`
	_, err = exec(dao.db, q, time.Now().UnixMilli(), name)
	if err != nil {
		return err
	}

	return nil
}

func (dao *ActivityDao) insert(a *model.Activity) error {
	q := `
		INSERT INTO activities(name, process, created_at, deleted_at)
		VALUES(?, ?, ?, -1);
	`
	_, err := exec(dao.db, q, a.Name, a.Process, time.Now().UnixMilli())
	if err != nil {
		return err
	}

	return nil
}

func (dao *ActivityDao) scan(rows *sql.Rows) (*model.Activity, error) {
	var a model.Activity
	err := rows.Scan(&a.Name, &a.Process, &a.CreatedAt, &a.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
