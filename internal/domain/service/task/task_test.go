// Package task focus on the data flow controlled by service layer
package task

import (
	"dddapib/internal/domain/model/aperr"
	"dddapib/internal/domain/model/entity"
	"dddapib/internal/infrastructure/persistence"
	"dddapib/internal/infrastructure/persistence/errors"
	mocktaskrepo "dddapib/internal/infrastructure/persistence/task/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaskService(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocktaskrepo.NewMockRepository(ctrl)
	svc := NewService(&persistence.Persistence{TaskRepository: mockRepo})

	t.Run("create-task", func(t *testing.T) {

		var taskToRepo *entity.Task
		mockRepo.EXPECT().Create(gomock.Any()).DoAndReturn(func(task *entity.Task) error {
			taskToRepo = task
			return nil
		})

		outputTask, err := svc.CreateTask("t1")
		assert.NoError(t, err)
		assert.Equal(t, taskToRepo, outputTask)

	})

	t.Run("list-tasks", func(t *testing.T) {

		repoOutput := []*entity.Task{entity.NewTask("t1")}
		mockRepo.EXPECT().
			List().
			Return(repoOutput, nil)

		tasks, err := svc.ListTasks()
		assert.NoError(t, err)
		assert.Equal(t, repoOutput, tasks)

	})

	t.Run("delete-task", func(t *testing.T) {

		taskID := "t1"
		mockRepo.EXPECT().Delete(taskID).Return(nil)

		err := svc.DeleteTask(taskID)
		assert.NoError(t, err)

	})

	t.Run("close-task-ok", func(t *testing.T) {

		var newStatus entity.TaskStatus

		t1 := entity.NewTask("t1")
		mockRepo.EXPECT().Get("t1").Return(t1, nil)
		mockRepo.EXPECT().
			UpdateStatus("t1", gomock.Any()).
			DoAndReturn(func(id string, status entity.TaskStatus) (*entity.Task, error) {
				newStatus = status
				return nil, nil
			})

		outputTask, err := svc.CloseTask("t1")
		assert.NoError(t, err)
		assert.Equal(t, outputTask, t1, "output task is not equal to the one from repo")
		assert.Equal(t, t1.Status, newStatus, "new status never been sent to repo")

	})

	t.Run("close-task-not-found", func(t *testing.T) {

		mockRepo.EXPECT().Get("t1").Return(nil, errors.ErrNotFound)
		_, err := svc.CloseTask("t1")
		assert.Equal(t, err, aperr.TaskNotFound)

	})

}

//
//func TestService_DeleteTask(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockRepo := mocktask.NewMockRepository(ctrl)
//	service := NewService(&persistence.Persistence{TaskRepository: mockRepo})
//
//	taskID := "task1"
//
//	mockRepo.EXPECT().Delete(taskID).Return(nil)
//
//	err := service.DeleteTask(taskID)
//	assert.NoError(t, err)
//}
//
//func TestService_CloseTask(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockRepo := mocktask.NewMockRepository(ctrl)
//	service := NewService(&persistence.Persistence{TaskRepository: mockRepo})
//
//	taskID := "task1"
//	task := entity.NewTask(taskID)
//	task.Status = entity.TaskStatusOpen
//
//	closedTask := *task
//	closedTask.Status = entity.TaskStatusDone
//
//	mockRepo.EXPECT().Get(taskID).Return(task, nil)
//	mockRepo.EXPECT().UpdateStatus(taskID, entity.TaskStatusDone).Return(&closedTask, nil)
//
//	result, err := service.CloseTask(taskID)
//	assert.NoError(t, err)
//	assert.Equal(t, &closedTask, result)
//}
//
//func TestService_CloseTask_NotFound(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockRepo := mocktask.NewMockRepository(ctrl)
//	service := NewService(&persistence.Persistence{TaskRepository: mockRepo})
//
//	taskID := "not-exist"
//
//	mockRepo.EXPECT().Get(taskID).Return(nil, errors.ErrNotFound)
//
//	_, err := service.CloseTask(taskID)
//	assert.Error(t, err)
//	assert.Equal(t, aperr.TaskNotFound, err)
//}
