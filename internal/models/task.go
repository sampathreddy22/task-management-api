package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Title       string    `gorm:"not null"`
	Description string
	Status      string //"todo", "in progress", "done"
	Priority    int
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uuid.UUID    `gorm:"type:uuid"` // Foreign key
	Comments    []Comment    `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE"`
	Attachments []Attachment `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE"`
}

func (t *Task) AfterFind(tx *gorm.DB) (err error) {

	if t.Comments == nil {
		if err := tx.Model(t).Association("Comments").Find(&t.Comments); err != nil {
			return err
		}
	}

	if t.Attachments == nil {
		if err := tx.Model(t).Association("Attachments").Find(&t.Attachments); err != nil {
			return err
		}
	}

	return nil
}
