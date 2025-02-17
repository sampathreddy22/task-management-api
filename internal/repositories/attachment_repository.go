package repositories

import (
	"github.com/google/uuid"
	"github.com/sampathreddy22/task-management-api/internal/models"
	"gorm.io/gorm"
)

type AttachmentRepository struct {
	*baseRepository[models.Attachment]
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) *AttachmentRepository {
	return &AttachmentRepository{
		baseRepository: NewBaseRepository[models.Attachment](db).(*baseRepository[models.Attachment]),
		db:             db,
	}
}

func (r *AttachmentRepository) Create(attachment *models.Attachment) error {
	return r.db.Create(attachment).Error
}

func (r *AttachmentRepository) GetByID(id uuid.UUID) (*models.Attachment, error) {
	var attachment models.Attachment
	err := r.db.First(&attachment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &attachment, nil
}

func (r *AttachmentRepository) GetByTaskID(taskID uuid.UUID) ([]models.Attachment, error) {
	var attachments []models.Attachment
	err := r.db.Where("task_id = ?", taskID).Find(&attachments).Error
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

func (r *AttachmentRepository) Update(attachment *models.Attachment) error {
	return r.db.Save(attachment).Error
}

func (r *AttachmentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Attachment{}, "id = ?", id).Error
}
