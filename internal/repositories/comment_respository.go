package repositories

import (
	"context"

	"github.com/sampathreddy22/task-management-api/internal/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	BaseRepository[models.Comment]
	GetByTaskID(ctx context.Context, taskID string) ([]models.Comment, error)
}

type commentRepository struct {
	*baseRepository[models.Comment]
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		baseRepository: NewBaseRepository[models.Comment](db).(*baseRepository[models.Comment]),
		db:             db,
	}
}

func (r *commentRepository) GetByTaskID(ctx context.Context, taskID string) ([]models.Comment, error) {
	var comments []models.Comment
	if err := r.db.Where("task_id = ?", taskID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
