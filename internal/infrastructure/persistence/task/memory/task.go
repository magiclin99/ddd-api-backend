package memory

import (
	"dddapib/internal/domain/model/entity"
	"dddapib/internal/infrastructure/persistence/task"
	"github.com/google/uuid"
)

type Repository struct {
	tasks map[uuid.UUID]*entity.Task
}

func (it *Repository) Create(task *entity.Task) error {
	it.tasks[task.ID] = task
	return nil
}
func (it *Repository) List() ([]*entity.Task, error) {
	output := make([]*entity.Task, 0, len(it.tasks))
	for _, t := range it.tasks {
		output = append(output, t)
	}
	return output, nil
}

func NewMemoryTaskRepository() task.Repository {
	return &Repository{
		tasks: map[uuid.UUID]*entity.Task{},
	}
}
