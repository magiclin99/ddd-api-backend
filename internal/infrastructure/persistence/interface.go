package persistence

import (
	"dddapib/internal/infrastructure/persistence/task"
	"dddapib/internal/infrastructure/persistence/task/memory"
)

type Persistence struct {
	TaskRepository task.Repository
}

func NewPersistence() *Persistence {
	return &Persistence{
		TaskRepository: memory.NewMemoryTaskRepository(),
	}
}
