package models

type CreateShortLinkRequest struct {
	OriginURL string `json:"origin_url"`
}

type CreateShortLinkResponse struct {
	ShortLinkID int64  `json:"id"`
	ShortLink   string `json:"short_link"`
	OriginURL   string `json:"origin_url"`
}
