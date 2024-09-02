package memory

import (
	"dddapib/internal/domain/model/entity"
	"dddapib/internal/infrastructure/persistence/errors"
	"dddapib/internal/infrastructure/persistence/task"
	"sync"
)

type Repository struct {
	tasks *syncMap[string, entity.Task]
}

func (it *Repository) Get(id string) (*entity.Task, error) {
	t, ok := it.tasks.Load(id)
	if !ok {
		return nil, errors.ErrNotFound
	}
	return t, nil
}

func (it *Repository) Create(task *entity.Task) error {
	_, load := it.tasks.LoadOrStore(task.ID, task)
	if load {
		return errors.ErrDuplicate
	}
	return nil
}
func (it *Repository) List() ([]*entity.Task, error) {
	return it.tasks.ToList(), nil
}

func (it *Repository) Delete(id string) error {
	// no existence check, even if the task does not exist, the result is the same.
	it.tasks.Delete(id)
	return nil
}

func (it *Repository) UpdateStatus(id string, status entity.TaskStatus) (*entity.Task, error) {
	t, ok := it.tasks.Load(id)
	if !ok {
		return nil, errors.ErrNotFound
	}
	t.Status = status
	return t, nil
}

func NewMemoryTaskRepository() task.Repository {
	return &Repository{
		tasks: &syncMap[string, entity.Task]{},
	}
}

type syncMap[K any, V any] struct {
	tasks sync.Map
}

func (it *syncMap[K, V]) LoadOrStore(key K, value *V) (*V, bool) {
	actual, loaded := it.tasks.LoadOrStore(key, value)
	return actual.(*V), loaded
}

func (it *syncMap[K, V]) Load(key K) (*V, bool) {
	actual, loaded := it.tasks.Load(key)
	if loaded {
		return actual.(*V), loaded
	} else {
		return nil, false
	}
}

func (it *syncMap[K, V]) Delete(key K) {
	it.tasks.Delete(key)
}

func (it *syncMap[K, V]) ToList() []*V {
	var output []*V
	it.tasks.Range(func(key, value interface{}) bool {
		output = append(output, value.(*V))
		return true
	})
	return output
}
