package task

import (
	"dddapib/internal/domain/model/entity"
)

type Repository interface {
	Create(task *entity.Task) error
	List() ([]*entity.Task, error)
	Delete(id string) error
}
