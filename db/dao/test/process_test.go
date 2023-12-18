package test

import (
	"fmt"
	"testing"

	"assalielmehdi/eventify/db/dao"
	"assalielmehdi/eventify/db/model"
)

func TestProcessDao(t *testing.T) {
	assert, db, teardown, err := setup(t)
	defer teardown()
	assert.Nil(err)

	dao := dao.NewProcessDao(db)
	generateProcesses(dao, 100)

	ps, err := dao.GetAll()
	assert.Nil(err)
	assert.Len(ps, 100)

	p, err := dao.GetByName("process-0")
	assert.Nil(err)
	assert.NotNil(p)
	assert.Equal(p.Name, "process-0")

	p, err = dao.GetByName("process-100")
	assert.NotNil(err)
	assert.Nil(p)

	err = dao.DeleteByName("process-1")
	assert.Nil(err)

	ps, err = dao.GetAll()
	assert.Nil(err)
	assert.Len(ps, 99)

	p, err = dao.GetByName("process-1")
	assert.NotNil(err)
	assert.Nil(p)
}

func generateProcesses(processDao *dao.ProcessDao, n int) {
	ps := make([]*model.Process, 0, n)

	for i := 0; i < n; i++ {
		ps = append(ps, &model.Process{
			Name: fmt.Sprintf("process-%d", i),
		})
		processDao.Save(ps[i])
	}
}
