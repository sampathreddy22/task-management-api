package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Title       string    `gorm:"not null"`
	Description string
	Status      string //"todo", "in progress", "done"
	Prirority   int
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uuid.UUID // Foreign key
	Comments    []Comment
	Attachments []Attachment
}
