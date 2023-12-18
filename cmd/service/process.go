package service

import (
	"database/sql"

	"github.com/eventified/eventified/db/dao"
	"github.com/eventified/eventified/db/model"
)

func GetProcessAll(db *sql.DB) ([]*model.Process, error) {
	return dao.GetProcessAll(db)
}

func GetProcessByName(db *sql.DB, name string) (*model.Process, error) {
	return dao.GetProcessByName(db, name)
}

func AddProcess(db *sql.DB, p *model.Process) error {
	err := validateProcess(p)
	if err != nil {
		return err
	}

	return dao.SaveProcess(db, p)
}

func DeleteProcessByName(db *sql.DB, name string) error {
	return dao.DeleteProcessByName(db, name)
}

func validateProcess(p *model.Process) error {
	errors := make([]string, 0)

	if len(p.Name) == 0 || len(p.Name) > 64 {
		errors = append(errors, "Name should have length n: 1 <= n <= 64")
	}

	if len(errors) == 0 {
		return nil
	}

	return ValidationError{errors}
}
