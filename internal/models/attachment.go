package models

import (
	"time"

	"github.com/google/uuid"
)

type Attachment struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	FileName   string    `gorm:"not null"`
	FilePath   string    `gorm:"not null"` //s3 URL or local path
	UploadedAt time.Time
	TaskID     uuid.UUID `gorm:"type:uuid"`
}

type AttachmentInput struct {
	FileName string    `json:"fileName" binding:"required"`
	FilePath string    `json:"filePath" binding:"required"`
	TaskID   uuid.UUID `json:"taskId" binding:"required"`
}

func (a *Attachment) FromAttachmentInput(input AttachmentInput) {
	a.ID = uuid.New()
	a.FileName = input.FileName
	a.FilePath = input.FilePath
	a.TaskID = input.TaskID
	a.UploadedAt = time.Now()
}

func (a *Attachment) TableName() string {
	return "attachments"
}
