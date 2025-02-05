package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID
	Content   string
	CreatedAt time.Time
	TaskID    uuid.UUID
	UserID    uuid.UUID
}
