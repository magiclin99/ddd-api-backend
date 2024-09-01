package memory

import (
	"dddapib/internal/domain/model/entity"
	"dddapib/internal/infrastructure/persistence/errors"
	"dddapib/internal/infrastructure/persistence/task"
)

type Repository struct {
	tasks map[string]*entity.Task
}

func (it *Repository) Get(id string) (*entity.Task, error) {
	t, ok := it.tasks[id]
	if !ok {
		return nil, errors.ErrNotFound
	}
	return t, nil
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

func (it *Repository) Delete(id string) error {
	// no existence check, even if the task does not exist, the result is the same.
	delete(it.tasks, id)
	return nil
}

func (it *Repository) UpdateStatus(id string, status entity.TaskStatus) (*entity.Task, error) {
	t, ok := it.tasks[id]
	if !ok {
		return nil, errors.ErrNotFound
	}
	t.Status = status
	return t, nil
}

func NewMemoryTaskRepository() task.Repository {
	return &Repository{
		tasks: map[string]*entity.Task{},
	}
}
