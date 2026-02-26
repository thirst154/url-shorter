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

func CreateURL(code string, originalURL string, expiresAt *time.Time) (*URL, error) {
	url := URL{
		Code:        code,
		OriginalURL: originalURL,
		ExpiresAt:   expiresAt,
	}
	return &url, DB.Create(&url).Error
}

func GetURL(code string) (*URL, error) {
	var url URL
	err := DB.First(&url, "code = ?", code).Error
	return &url, err
}

func IsCodeUnique(code string) bool {
	var url URL
	err := DB.First(&url, "code = ?", code).Error
	return err == gorm.ErrRecordNotFound
}

func IncrementClicks(id uint) {
	DB.Model(&URL{}).Where("id = ?", id).UpdateColumn("clicks", gorm.Expr("clicks + 1"))
}
