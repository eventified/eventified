package test

import (
	"database/sql"
	"fmt"
	"net/http"
	"testing"

	"github.com/eventified/eventified/common"
	"github.com/eventified/eventified/db/dao"
	"github.com/eventified/eventified/db/model"
)

func TestProcessDao(t *testing.T) {
	assert, db, teardown, err := setup(t)
	defer teardown()
	assert.Nil(err)

	generateProcesses(db, 100)

	ps, err := dao.GetProcessAll(db)
	assert.Nil(err)
	assert.Len(ps, 100)

	p, err := dao.GetProcessByName(db, "process-0")
	assert.Nil(err)
	assert.NotNil(p)
	assert.Equal(p.Name, "process-0")

	p, err = dao.GetProcessByName(db, "process-100")
	assert.NotNil(err)
	assert.Equal(err.(*common.Error).Status, http.StatusNotFound)
	assert.Nil(p)

	err = dao.DeleteProcessByName(db, "process-1")
	assert.Nil(err)

	ps, err = dao.GetProcessAll(db)
	assert.Nil(err)
	assert.Len(ps, 99)

	p, err = dao.GetProcessByName(db, "process-1")
	assert.NotNil(err)
	assert.Equal(err.(*common.Error).Status, http.StatusNotFound)
	assert.Nil(p)
}

func generateProcesses(db *sql.DB, n int) {
	ps := make([]*model.Process, 0, n)

	for i := 0; i < n; i++ {
		ps = append(ps, &model.Process{
			Name: fmt.Sprintf("process-%d", i),
		})
		dao.SaveProcess(db, ps[i])
	}
}
