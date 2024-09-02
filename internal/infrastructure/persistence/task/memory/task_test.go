package memory

import (
	"dddapib/internal/domain/model/entity"
	"dddapib/internal/infrastructure/persistence/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaskMemoryRepository(t *testing.T) {

	repo := NewMemoryTaskRepository()

	// create task t1
	t1 := entity.NewTask("t1")
	err := repo.Create(t1)
	assert.NoError(t, err)

	// do not allow duplicate task
	err = repo.Create(t1)
	assert.Equal(t, err, errors.ErrDuplicate)

	// create task t2
	t2 := entity.NewTask("t2")
	err = repo.Create(t2)
	assert.NoError(t, err)

	// get task t1
	taskFromRepo, err := repo.Get(t1.ID)
	assert.NoError(t, err)
	assert.Equal(t, t1, taskFromRepo)

	// can't get task that does not exist
	_, err = repo.Get("not-exist")
	assert.Equal(t, errors.ErrNotFound, err)

	// list tasks, should be t1, t2
	taskList, err := repo.List()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(taskList))
	assert.Contains(t, taskList, t1)
	assert.Contains(t, taskList, t2)

	// update t1 status to done
	updatedTask, err := repo.UpdateStatus(t1.ID, entity.TaskStatusDone)
	assert.NoError(t, err)
	assert.Equal(t, entity.TaskStatusDone, updatedTask.Status)
	assert.Equal(t, t1.ID, updatedTask.ID)

	// can't update task that does not exist
	_, err = repo.UpdateStatus("not-exist", entity.TaskStatusDone)
	assert.Equal(t, errors.ErrNotFound, err)

	// delete all tasks
	err = repo.Delete(t1.ID)
	assert.NoError(t, err)
	err = repo.Delete(t2.ID)
	assert.NoError(t, err)
	taskList, err = repo.List()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(taskList))

}
