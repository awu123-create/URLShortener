package Models

import (
	"time"
)

type URL struct {
	ID          uint   `gorm:"primaryKey"`
	OriginalURL string `gorm:"type:text;not null"`
	ShortCode   string `gorm:"size:20;uniqueIndex;not null"`
	CreatedAt   time.Time
	ExpiresAt   *time.Time
	VisitCount  int `gorm:"default:0"`
}
