package service

import (
	"database/sql"

	"github.com/eventified/eventified/common"
	"github.com/eventified/eventified/db/dao"
	"github.com/eventified/eventified/db/model"
)

func GetActivityAll(db *sql.DB) ([]*model.Activity, *common.Error) {
	return dao.GetActivityAll(db)
}

func GetActivityByName(db *sql.DB, name string) (*model.Activity, *common.Error) {
	return dao.GetActivityByName(db, name)
}

func AddActivity(db *sql.DB, a *model.Activity) *common.Error {
	err := validateActivity(db, a)
	if err != nil {
		return err
	}

	return dao.SaveActivity(db, a)
}

func DeleteActivityByName(db *sql.DB, name string) *common.Error {
	return dao.DeleteActivityByName(db, name)
}

func validateActivity(db *sql.DB, a *model.Activity) *common.Error {
	_, err := dao.GetProcessByName(db, a.Process)
	if err != nil {
		return err
	}

	if len(a.Name) == 0 || len(a.Name) > 64 {
		return common.BadRequestError("Name should have length n: 1 <= n <= 64")
	}

	return nil
}
