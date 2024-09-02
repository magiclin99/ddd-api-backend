package task

import (
	"context"
	"dddapib/internal/domain/service"
	"dddapib/internal/domain/service/task"
	"dddapib/internal/infrastructure/transport/http/dto"
	"dddapib/internal/infrastructure/transport/http/handler"
	"github.com/gin-gonic/gin"
)

var taskService task.Service

func Init(router *gin.Engine, svc *service.Service) {

	taskService = svc.TaskService

	// register API handler
	taskApiGroup := router.Group("tasks")
	taskApiGroup.GET("", handler.JsonWithoutPayload(listTasks))
	taskApiGroup.POST("", handler.Json[dto.CreateTaskRequest](createTask))
	taskApiGroup.DELETE(":taskID", handler.JsonWithoutPayload(deleteTask))
	taskApiGroup.PUT(":taskID", handler.JsonWithoutPayload(updateTask))
}

func listTasks(ctx context.Context, g *gin.Context) (any, error) {
	// simply return the list of entities for now.
	// add dto mapping if needed
	return taskService.ListTasks()
}

func createTask(ctx context.Context, g *gin.Context, payload *dto.CreateTaskRequest) (any, error) {
	return taskService.CreateTask(payload.Name)
}

func deleteTask(ctx context.Context, g *gin.Context) (any, error) {
	taskID := g.Param("taskID")
	err := taskService.DeleteTask(taskID)
	return nil, err
}

func updateTask(ctx context.Context, g *gin.Context) (any, error) {
	taskID := g.Param("taskID")
	return taskService.CloseTask(taskID)
}
