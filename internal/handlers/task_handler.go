package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sampathreddy22/task-management-api/internal/models"
	"github.com/sampathreddy22/task-management-api/internal/services"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

// @Summary Create a new task
// @Description Create a new task with the given title and description
// @Accept json
// @Produce json
// @Param task body models.Task true "Task to create"
// @Success 201 {object} models.Task
// @Router /tasks [post]
// @tags tasks
func (h *TaskHandler) CreateTask(c *gin.Context) {
	// get the task from the request body and create a new task using the task service
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID = uuid.New()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	if err := h.taskService.CreateTask(context.Background(), &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// @Summary Get a task by ID
// @Description Get a task by ID
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} models.Task
// @Router /tasks/{id} [get]
// @tags tasks
func (h *TaskHandler) GetTaskByID(c *gin.Context) {

}

// @Summary Update a task
// @Description Update a task by ID
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body models.Task true "Task to update"
// @Success 200 {object} models.Task
// @Router /tasks/{id} [put]
// @tags tasks
func (h *TaskHandler) UpdateTask(c *gin.Context) {

}

// @Summary Delete a task
// @Description Delete a task by ID
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 204
// @Router /tasks/{id} [delete]
// @tags tasks
func (h *TaskHandler) DeleteTask(c *gin.Context) {

}

// @Summary Get all tasks
// @Description Get all tasks
// @Accept json
// @Produce json
// @Success 200 {array} models.Task
// @Router /tasks [get]
// @tags tasks
func (h *TaskHandler) GetTasks(c *gin.Context) {

}
