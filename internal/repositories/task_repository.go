package repositories

import (
	"context"
	"strings"

	"github.com/sampathreddy22/task-management-api/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	BaseRepository[models.Task]
	GetByUserID(ctx context.Context, userID string, offset, limit int) ([]models.Task, error)
	GetByStatus(ctx context.Context, status string, offset, limit int) ([]models.Task, error)
	GetByPriority(ctx context.Context, priority string, offset, limit int) ([]models.Task, error)
	search(
		ctx context.Context,
		query string,
		offset, limit int,
	) ([]models.Task, error)
	GetAll(ctx context.Context, offset, limit int) ([]models.Task, error)
}

type taskRepository struct {
	*baseRepository[models.Task]
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		baseRepository: NewBaseRepository[models.Task](db).(*baseRepository[models.Task]),
		db:             db,
	}
}

func (r *taskRepository) GetByUserID(ctx context.Context, userID string, offset, limit int) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.WithContext(ctx).Where("user_id=?", userID).Offset(offset).Limit(limit).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) GetByStatus(ctx context.Context, status string, offset, limit int) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.WithContext(ctx).Where("status=?", status).Offset(offset).Limit(limit).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) GetByPriority(ctx context.Context, priority string, offset, limit int) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.WithContext(ctx).Where("priority=?", priority).Offset(offset).Limit(limit).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) search(ctx context.Context, query string, offset, limit int) ([]models.Task, error) {
	var tasks []models.Task
	searchPattern := "%" + strings.ReplaceAll(strings.ReplaceAll(query, "%", "\\%"), "_", "\\_") + "%"
	if err := r.db.WithContext(ctx).
		Where("LOWER(title) LIKE LOWER(?) OR LOWER(description) LIKE LOWER(?)",
			searchPattern,
			searchPattern).
		Offset(offset).
		Limit(limit).
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) GetAll(ctx context.Context, offset, limit int) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
