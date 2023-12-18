package service

import (
	"database/sql"

	"github.com/eventified/eventified/common"
	"github.com/eventified/eventified/db/dao"
	"github.com/eventified/eventified/db/model"
)

func GetProcessAll(db *sql.DB) ([]*model.Process, *common.Error) {
	return dao.GetProcessAll(db)
}

func GetProcessByName(db *sql.DB, name string) (*model.Process, *common.Error) {
	return dao.GetProcessByName(db, name)
}

func AddProcess(db *sql.DB, p *model.Process) *common.Error {
	err := validateProcess(p)
	if err != nil {
		return err
	}

	return dao.SaveProcess(db, p)
}

func DeleteProcessByName(db *sql.DB, name string) *common.Error {
	return dao.DeleteProcessByName(db, name)
}

func validateProcess(p *model.Process) *common.Error {
	if len(p.Name) == 0 || len(p.Name) > 64 {
		return common.BadRequestError("Name should have length n: 1 <= n <= 64")
	}

	return nil
}
