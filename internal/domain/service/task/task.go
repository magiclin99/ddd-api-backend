// Package task holds all the domain logic for the task domain.
package task

import (
	"dddapib/internal/domain/model/entity"
	"dddapib/internal/infrastructure/persistence"
	"dddapib/internal/infrastructure/persistence/task"
)

func NewService(p *persistence.Persistence) Service {
	return &serviceImpl{
		taskRepo: p.TaskRepository,
	}
}

type Service interface {
	CreateTask(name string) error
	ListTasks() ([]*entity.Task, error)
}

type serviceImpl struct {
	taskRepo task.Repository
}

func (it *serviceImpl) CreateTask(name string) error {
	newTask := entity.NewTask(name)
	err := it.taskRepo.Create(newTask)
	return err
}

func (it *serviceImpl) ListTasks() ([]*entity.Task, error) {
	return it.taskRepo.List()
}
