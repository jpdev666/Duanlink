package db_models

import (
	"time"
)

type ShortLink struct {
	ID        int64  `gorm:"primaryKey"`
	ShortCode string `gorm:"index"`
	OriginURL string `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *ShortLink) TableName() string {
	return "short_links"
}
