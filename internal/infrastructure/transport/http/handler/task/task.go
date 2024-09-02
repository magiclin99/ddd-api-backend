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

// listTasks handles the GET request to list all tasks.
//
//	@Summary	List all tasks
//	@Tags		tasks
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}	entity.Task
//	@Router		/tasks [get]
func listTasks(ctx context.Context, g *gin.Context) (any, error) {
	// simply return the list of entities for now.
	// add dto mapping if needed
	return taskService.ListTasks()
}

// createTask handles the POST request to create a new task.
//
//	@Summary	Create a new task
//	@Tags		tasks
//	@Accept		json
//	@Produce	json
//	@Param		payload	body		dto.CreateTaskRequest	true	"Task creation request"
//	@Success	200		{object}	entity.Task
//	@Failure	400		{object}	dto.ApiError
//
//	@Router		/tasks [post]
func createTask(ctx context.Context, g *gin.Context, payload *dto.CreateTaskRequest) (any, error) {
	return taskService.CreateTask(payload.Name)
}

// deleteTask handles the DELETE request to delete a task by its ID.
//
//	@Summary	Delete a task
//	@Tags		tasks
//	@Accept		json
//	@Produce	json
//	@Param		taskID	path	string	true	"Task ID"
//	@Success	200		"Task successfully deleted (even the task does not exist)"
//	@Router		/tasks/{taskID} [delete]
func deleteTask(ctx context.Context, g *gin.Context) (any, error) {
	taskID := g.Param("taskID")
	err := taskService.DeleteTask(taskID)
	return nil, err
}

// updateTask handles the PUT request to update a task by its ID.
//
//	@Summary		Close a task immediately
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			taskID	path		string			true	"Task ID"
//	@Success		200		{object}	entity.Task		"Task successfully updated"
//	@Failure		400		{object}	dto.ApiError	"TASK-00001 - Task not found"
//	@Router			/tasks/{taskID} [put]
//	@Description	Close a task immediately
func updateTask(ctx context.Context, g *gin.Context) (any, error) {
	taskID := g.Param("taskID")
	return taskService.CloseTask(taskID)
}
