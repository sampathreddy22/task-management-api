package services

import (
	"github.com/google/uuid"
	"github.com/sampathreddy22/task-management-api/internal/models"
	"github.com/sampathreddy22/task-management-api/internal/repositories"
)

type AttachmentService struct {
	repo *repositories.AttachmentRepository
}

func NewAttachmentService(repo *repositories.AttachmentRepository) *AttachmentService {
	return &AttachmentService{repo: repo}
}

func (s *AttachmentService) CreateAttachment(input models.AttachmentInput) (*models.Attachment, error) {
	attachment := &models.Attachment{}
	attachment.FromAttachmentInput(input)

	err := s.repo.Create(attachment)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

func (s *AttachmentService) GetAttachment(id uuid.UUID) (*models.Attachment, error) {
	return s.repo.GetByID(id)
}

func (s *AttachmentService) GetTaskAttachments(taskID uuid.UUID) ([]models.Attachment, error) {
	return s.repo.GetByTaskID(taskID)
}

func (s *AttachmentService) UpdateAttachment(id uuid.UUID, input models.AttachmentInput) (*models.Attachment, error) {
	attachment, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	attachment.FileName = input.FileName
	attachment.FilePath = input.FilePath

	err = s.repo.Update(attachment)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

func (s *AttachmentService) DeleteAttachment(id uuid.UUID) error {
	return s.repo.Delete(id)
}
