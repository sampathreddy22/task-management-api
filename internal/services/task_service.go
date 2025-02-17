package services

import (
	"context"

	"github.com/sampathreddy22/task-management-api/internal/models"
	"github.com/sampathreddy22/task-management-api/internal/repositories"
)

type TaskService struct {
	taskRepo repositories.TaskRepository
}

func NewTaskService(taskRepo repositories.TaskRepository) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, task *models.Task) error {
	return s.taskRepo.Create(ctx, task)
}

func (s *TaskService) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	return s.taskRepo.GetByID(ctx, id)
}

func (s *TaskService) UpdateTask(ctx context.Context, task *models.Task) error {
	return s.taskRepo.Update(ctx, task)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.taskRepo.Delete(ctx, id)
}

func (s *TaskService) GetTasksByUserID(ctx context.Context, userID string, offset, limit int) ([]models.Task, error) {
	return s.taskRepo.GetByUserID(ctx, userID, offset, limit)
}

func (s *TaskService) GetTasksByStatus(ctx context.Context, status string, offset, limit int) ([]models.Task, error) {
	return s.taskRepo.GetByStatus(ctx, status, offset, limit)
}

func (s *TaskService) GetTasks(ctx context.Context, offset, limit int) ([]models.Task, error) {
	return s.taskRepo.GetAll(ctx, offset, limit)
}
