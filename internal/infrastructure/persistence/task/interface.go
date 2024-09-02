package task

import (
	"dddapib/internal/domain/model/entity"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen -destination=mock/mock_task_repo.go -source=interface.go --package=mocktaskrepo
type Repository interface {
	Get(id string) (*entity.Task, error)
	Create(task *entity.Task) error
	List() ([]*entity.Task, error)
	Delete(id string) error
	UpdateStatus(id string, status entity.TaskStatus) (*entity.Task, error)
}
