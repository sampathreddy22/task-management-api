package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sampathreddy22/task-management-api/internal/services"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

// CreateTask creates a new task
func (h *TaskHandler) CreateTask(c *gin.Context) {
	// get the task from the request body and create a new task using the task service

}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {

}

func (h *TaskHandler) UpdateTask(c *gin.Context) {

}

func (h *TaskHandler) DeleteTask(c *gin.Context) {

}

func (h *TaskHandler) GetTasks(c *gin.Context) {

}
