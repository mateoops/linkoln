package services

import (
	"context"
	"mateoops/linkoln/models"
	"mateoops/linkoln/repositories"
)

type ShortService struct {
	repo repositories.ShortRepo
}

func NewShortService(sr repositories.ShortRepo) *ShortService {
	return &ShortService{
		repo: sr,
	}
}

func (ss *ShortService) CreateShort(ctx context.Context, short models.Short) (string, error) {
	return ss.repo.CreateShort(ctx, short)
}

func (ss *ShortService) GetByShortUrl(ctx context.Context, shortUrl string) models.Short {
	return ss.repo.GetByShortUrl(ctx, shortUrl)
}
