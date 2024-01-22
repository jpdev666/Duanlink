package services

import (
	"fmt"
	"time"

	"github.com/ved2pj/Duanlink/internal/models"
	"github.com/ved2pj/Duanlink/internal/repos"
	"github.com/ved2pj/Duanlink/internal/repos/db_models"
	"github.com/ved2pj/Duanlink/internal/utils"
)

type ShortLinkService interface {
	Create(req models.CreateShortLinkRequest) (*models.CreateShortLinkResponse, error)
	LookupByShortCode(shortCode string) (*db_models.ShortLink, error)
}

type shortLinkService struct {
	repo repos.ShortLinkRepo
}

func NewShortLinkService(repo repos.ShortLinkRepo) ShortLinkService {
	return &shortLinkService{repo: repo}
}

func (s *shortLinkService) Create(req models.CreateShortLinkRequest) (*models.CreateShortLinkResponse, error) {
	shortLink := &db_models.ShortLink{
		ShortCode: utils.GenerateShortCode(),
		OriginURL: req.OriginURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(shortLink); err != nil {
		return nil, err
	}
	return &models.CreateShortLinkResponse{
		ShortLinkID: shortLink.ID,
		ShortLink:   fmt.Sprintf("http://127.0.0.1/%s", shortLink.ShortCode),
		OriginURL:   req.OriginURL,
	}, nil
}

func (s *shortLinkService) LookupByShortCode(shortCode string) (*db_models.ShortLink, error) {
	return s.repo.LookupByShortCode(shortCode)
}
