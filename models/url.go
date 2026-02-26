package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	Code        string     `gorm:"uniqueIndex;not null;size:16"    json:"code"`
	OriginalURL string     `gorm:"not null"                        json:"original_url"`
	Clicks      uint       `gorm:"default:0"                       json:"clicks"`
	ExpiresAt   *time.Time `gorm:"index"                           json:"expires_at,omitempty"`
}
