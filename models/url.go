package models

import (
	"time"
)

type Url struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	OriginalURL string    `gorm:"type:text;not null" json:"original_url"`
	ShortCode   string    `gorm:"type:varchar(10);unique;not null" json:"short_code"`
	ClickCount  uint      `gorm:"default:0" json:"click_count"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}