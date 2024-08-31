// Package service exposes the service layer of the application
package service

import (
	"dddapib/internal/domain/service/task"
	"dddapib/internal/infrastructure/persistence"
)

type Service struct {
	TaskService task.Service
}

func NewService(p *persistence.Persistence) *Service {

	return &Service{
		TaskService: task.NewService(p),
		// add more service category here
	}
}
