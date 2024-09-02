package entity

import (
	"github.com/google/uuid"
)

type TaskStatus int

var (
	TaskStatusToDo TaskStatus = 0
	TaskStatusDone TaskStatus = 1
)

type Task struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Status TaskStatus `json:"status"`
}

func (t *Task) Close() TaskStatus {
	t.Status = TaskStatusDone
	return t.Status
}

// NewTask creates a new Task with initial status
func NewTask(name string) *Task {
	return &Task{
		ID:     "task-" + uuid.New().String(),
		Name:   name,
		Status: TaskStatusToDo,
	}
}
