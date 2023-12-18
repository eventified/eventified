package test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/eventified/eventified/db/dao"
	"github.com/eventified/eventified/db/model"
)

func TestActivityDao(t *testing.T) {
	assert, db, teardown, err := setup(t)
	defer teardown()
	assert.Nil(err)

	generateProcesses(db, 10)
	generateActivities(db, 100)

	as, err := dao.GetActivityAll(db)
	assert.Nil(err)
	assert.Len(as, 100)

	a, err := dao.GetActivityByName(db, "activity-0")
	assert.Nil(err)
	assert.NotNil(a)
	assert.Equal(a.Name, "activity-0")

	a, err = dao.GetActivityByName(db, "activity-100")
	assert.NotNil(err)
	assert.Nil(a)

	err = dao.DeleteActivityByName(db, "activity-1")
	assert.Nil(err)

	as, err = dao.GetActivityAll(db)
	assert.Nil(err)
	assert.Len(as, 99)

	a, err = dao.GetActivityByName(db, "activity-1")
	assert.NotNil(err)
	assert.Nil(a)
}

func generateActivities(db *sql.DB, n int) {
	as := make([]*model.Activity, 0, n)

	for i := 0; i < n; i++ {
		as = append(as, &model.Activity{
			Name:    fmt.Sprintf("activity-%d", i),
			Process: fmt.Sprintf("process-%d", i%10),
		})
		dao.SaveActivity(db, as[i])
	}
}
