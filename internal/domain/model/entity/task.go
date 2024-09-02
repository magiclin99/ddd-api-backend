package entity

import (
	"github.com/google/uuid"
)

type TaskStatus int

var (
	TaskStatusToDo TaskStatus = 0
	TaskStatusDone TaskStatus = 1
)

// Task represents a task in the system.
//
// @swagger:model Task
type Task struct {
	// The unique identifier of the task
	ID string `json:"id"`
	// Task name
	Name string `json:"name"`
	// 0 - ToDo, 1 - Done
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
