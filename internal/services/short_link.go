package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ved2pj/Duanlink/internal/datastore"
	"github.com/ved2pj/Duanlink/internal/models"
	"github.com/ved2pj/Duanlink/internal/repos"
	"github.com/ved2pj/Duanlink/internal/repos/db_models"
	"github.com/ved2pj/Duanlink/internal/utils"
)

const (
	defaultExpiration = 24 * time.Hour
)

type ShortLinkService interface {
	Create(req models.CreateShortLinkRequest) (*models.CreateShortLinkResponse, error)
	LookupByShortCode(shortCode string) (*db_models.ShortLink, error)
}

type shortLinkService struct {
	repo     repos.ShortLinkRepo
	redisCli *redis.Client
}

func NewShortLinkService(repo repos.ShortLinkRepo) ShortLinkService {
	return &shortLinkService{
		repo:     repo,
		redisCli: datastore.Get().Redis,
	}
}

func (s *shortLinkService) Create(req models.CreateShortLinkRequest) (*models.CreateShortLinkResponse, error) {
	shortLink := &db_models.ShortLink{
		ShortCode: utils.GenerateShortCode(),
		OriginURL: req.OriginURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ExpiredAt: req.ExpiredAt,
	}

	if err := s.repo.Create(shortLink); err != nil {
		return nil, err
	}
	return &models.CreateShortLinkResponse{
		ShortLinkID: shortLink.ID,
		ShortLink:   fmt.Sprintf("http://127.0.0.1:8080/%s", shortLink.ShortCode),
		OriginURL:   req.OriginURL,
	}, nil
}

func (s *shortLinkService) LookupByShortCode(shortCode string) (*db_models.ShortLink, error) {
	shortLink, err := s.lookupFromRedis(shortCode)
	if err != nil {
		return nil, err
	}
	if shortLink != nil {
		return shortLink, nil
	}

	shortLink, err = s.repo.LookupByShortCode(shortCode)
	if err != nil {
		return nil, err
	}
	if shortLink.ExpiredAt != nil && shortLink.ExpiredAt.Before(time.Now()) {
		return nil, models.ErrorShortLinkExpired
	}

	// Calculate the expiration time
	expiration := defaultExpiration
	if shortLink.ExpiredAt != nil {
		expiration = time.Until(*shortLink.ExpiredAt)
	}

	// Set the short link to Redis with the calculated expiration time
	if err = s.redisCli.Set(context.Background(), shortCode, shortLink.OriginURL, expiration).Err(); err != nil {
		return nil, err
	}

	return shortLink, nil
}

func (s *shortLinkService) lookupFromRedis(shortCode string) (*db_models.ShortLink, error) {
	val, err := s.redisCli.Get(context.Background(), shortCode).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	return &db_models.ShortLink{
		ShortCode: shortCode,
		OriginURL: val,
	}, nil
}
