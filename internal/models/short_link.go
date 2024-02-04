package models

import (
	"fmt"
	"time"
)

var (
	ErrorShortLinkExpired = fmt.Errorf("short link expired")
)

type CreateShortLinkRequest struct {
	OriginURL string     `json:"origin_url"`
	ExpiredAt *time.Time `json:"expired_at"`
}

type CreateShortLinkResponse struct {
	ShortLinkID int64  `json:"id"`
	ShortLink   string `json:"short_link"`
	OriginURL   string `json:"origin_url"`
}
