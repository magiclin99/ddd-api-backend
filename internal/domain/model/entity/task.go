package entity

import (
	"github.com/google/uuid"
)

type TaskStatus int

var (
	InProgressToDo TaskStatus = 0
	TaskStatusDone TaskStatus = 1
)

type Task struct {
	ID     uuid.UUID  `json:"id"`
	Name   string     `json:"name"`
	Status TaskStatus `json:"status"`
}

// NewTask creates a new Task with initial status
func NewTask(name string) *Task {

	switch {
	case len(name) == 0:
	case len(name) > 100:

	}

	return &Task{
		ID:     uuid.New(),
		Name:   name,
		Status: InProgressToDo,
	}
}
