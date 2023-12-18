package test

import (
	"fmt"
	"testing"

	"assalielmehdi/eventify/db/dao"
	"assalielmehdi/eventify/db/model"
)

func TestActivityDao(t *testing.T) {
	assert, db, teardown, err := setup(t)
	defer teardown()
	assert.Nil(err)

	processDao := dao.NewProcessDao(db)
	generateProcesses(processDao, 10)

	dao := dao.NewActivityDao(db)
	generateActivities(dao, 100)

	as, err := dao.GetAll()
	assert.Nil(err)
	assert.Len(as, 100)

	a, err := dao.GetByName("activity-0")
	assert.Nil(err)
	assert.NotNil(a)
	assert.Equal(a.Name, "activity-0")

	a, err = dao.GetByName("activity-100")
	assert.NotNil(err)
	assert.Nil(a)

	err = dao.DeleteByName("activity-1")
	assert.Nil(err)

	as, err = dao.GetAll()
	assert.Nil(err)
	assert.Len(as, 99)

	a, err = dao.GetByName("activity-1")
	assert.NotNil(err)
	assert.Nil(a)
}

func generateActivities(activityDao *dao.ActivityDao, n int) {
	as := make([]*model.Activity, 0, n)

	for i := 0; i < n; i++ {
		as = append(as, &model.Activity{
			Name:    fmt.Sprintf("activity-%d", i),
			Process: fmt.Sprintf("process-%d", i%10),
		})
		activityDao.Save(as[i])
	}
}
