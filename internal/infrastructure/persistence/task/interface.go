package task

import (
	"dddapib/internal/domain/model/entity"
)

type Repository interface {
	Get(id string) (*entity.Task, error)
	Create(task *entity.Task) error
	List() ([]*entity.Task, error)
	Delete(id string) error
	UpdateStatus(id string, status entity.TaskStatus) (*entity.Task, error)
}
