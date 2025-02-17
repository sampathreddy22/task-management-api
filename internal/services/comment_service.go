package services

import (
	"context"

	"github.com/sampathreddy22/task-management-api/internal/models"
	"github.com/sampathreddy22/task-management-api/internal/repositories"
)

type CommentService struct {
	commentRepo repositories.CommentRepository
}

func NewCommentService(commentRepo repositories.CommentRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
	}
}

func (s *CommentService) CreateComment(ctx context.Context, comment *models.Comment) error {
	return s.commentRepo.Create(ctx, comment)
}

func (s *CommentService) GetCommentByID(ctx context.Context, id string) (*models.Comment, error) {
	return s.commentRepo.GetByID(ctx, id)
}

func (s *CommentService) GetCommentsByTaskID(ctx context.Context, taskID string) ([]models.Comment, error) {
	return s.commentRepo.GetByTaskID(ctx, taskID)
}

func (s *CommentService) DeleteComment(ctx context.Context, id string) error {
	return s.commentRepo.Delete(ctx, id)
}
