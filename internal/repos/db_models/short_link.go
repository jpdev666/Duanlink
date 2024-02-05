package db_models

import (
	"time"
)

type ShortLink struct {
	ID        int64      `gorm:"primaryKey" json:"id,omitempty"`
	ShortCode string     `gorm:"index" json:"short_code,omitempty"`
	OriginURL string     `gorm:"index" json:"origin_url,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
}

func (s *ShortLink) TableName() string {
	return "short_links"
}
