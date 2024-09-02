// Package task holds all the domain logic for the task domain.
package task

import (
	"dddapib/internal/domain/model/aperr"
	"dddapib/internal/domain/model/entity"
	"dddapib/internal/infrastructure/persistence"
	persistenceErr "dddapib/internal/infrastructure/persistence/errors"
	"dddapib/internal/infrastructure/persistence/task"
	"errors"
)

func NewService(p *persistence.Persistence) Service {
	return &serviceImpl{
		taskRepo: p.TaskRepository,
	}
}

//go:generate go run -mod=mod github.com/golang/mock/mockgen -destination=mock/task.go -source=task.go --package=mocktasksvc
type Service interface {
	// CreateTask create a new task
	CreateTask(name string) (*entity.Task, error)
	// ListTasks return all tasks
	ListTasks() ([]*entity.Task, error)
	// DeleteTask remove task by id
	DeleteTask(id string) error
	// CloseTask move task to done
	CloseTask(id string) (*entity.Task, error)
}

type serviceImpl struct {
	taskRepo task.Repository
}

func (it *serviceImpl) CreateTask(name string) (*entity.Task, error) {
	newTask := entity.NewTask(name)
	err := it.taskRepo.Create(newTask)
	return newTask, err
}

func (it *serviceImpl) ListTasks() ([]*entity.Task, error) {
	return it.taskRepo.List()
}

func (it *serviceImpl) DeleteTask(id string) error {
	return it.taskRepo.Delete(id)
}

func (it *serviceImpl) CloseTask(id string) (*entity.Task, error) {
	taskToClose, err := it.taskRepo.Get(id)
	if err != nil {
		return nil, toApError(err)
	}

	newStatus := taskToClose.Close()

	_, err = it.taskRepo.UpdateStatus(id, newStatus)
	return taskToClose, toApError(err)
}

// translate persistence error to application error
func toApError(err error) error {
	switch {
	case errors.Is(err, persistenceErr.ErrNotFound):
		return aperr.TaskNotFound
	default:
		return err
	}
}
