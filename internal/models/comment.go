package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Content   string
	CreatedAt time.Time
	TaskID    uuid.UUID `gorm:"type:uuid"`
	UserID    uuid.UUID `gorm:"type:uuid"`
}

func (c *Comment) TableName() string {
	return "comments"
}
