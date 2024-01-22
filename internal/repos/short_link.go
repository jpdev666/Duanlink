package repos

import (
	"github.com/ved2pj/Duanlink/internal/repos/db_models"
	"gorm.io/gorm"
)

type ShortLinkRepo interface {
	Create(shortLink *db_models.ShortLink) error
	LookupByShortCode(shortCode string) (*db_models.ShortLink, error)
}

type shortLinkRepo struct {
	db *gorm.DB
}

func NewShortLinkRepo(db *gorm.DB) ShortLinkRepo {
	return &shortLinkRepo{db: db}
}

func (r *shortLinkRepo) Create(shortLink *db_models.ShortLink) error {
	return r.db.Create(shortLink).Error
}

func (r *shortLinkRepo) LookupByShortCode(shortCode string) (*db_models.ShortLink, error) {
	var shortLink *db_models.ShortLink
	tx := r.db.Find(&shortLink, "short_code = ?", shortCode)
	return shortLink, tx.Error
}
