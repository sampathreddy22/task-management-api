package models

import (
	"time"

	"github.com/google/uuid"
)

type Attachment struct {
	ID        uuid.UUID
	FileNmae  string
	FilePath  string //s3 URL or local path
	UplodedAt time.Time
	TaskID    uuid.UUID
}
