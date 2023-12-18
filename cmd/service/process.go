package service

import (
	"assalielmehdi/eventify/db/dao"
	"assalielmehdi/eventify/db/model"
)

type ProcessService struct {
	processDao *dao.ProcessDao
}

func (svc *ProcessService) GetAll() ([]*model.Process, error) {
	return svc.processDao.GetAll()
}

func (svc *ProcessService) GetByName(name string) (*model.Process, error) {
	return svc.processDao.GetByName(name)
}

func (svc *ProcessService) Add(p *model.Process) error {
	err := validate(p)
	if err != nil {
		return err
	}
	return svc.processDao.Save(p)
}

func (svc *ProcessService) DeleteByName(name string) error {
	return svc.processDao.DeleteByName(name)
}

func validate(p *model.Process) error {
	errors := make([]string, 0)

	if len(p.Name) == 0 || len(p.Name) > 64 {
		errors = append(errors, "Name should have length n: 1 <= n <= 64")
	}

	if len(errors) == 0 {
		return nil
	}

	return ValidationError{errors}
}
