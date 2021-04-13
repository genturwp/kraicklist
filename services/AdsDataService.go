package services

import (
	"context"

	"challenge.haraj.com.sa/kraicklist/entities"
	"challenge.haraj.com.sa/kraicklist/repositories"
)

const adsDataDir = "data-dir"

type AdsDataService interface {
	SearchAdsData(ctx context.Context, searchStr string) ([]*entities.AdsData, error)
}

type service struct {
	repo *repositories.Repository
}

func NewAdsDataService(repo *repositories.Repository) AdsDataService {
	return &service{
		repo: repo,
	}
}

func (service *service) SearchAdsData(ctx context.Context, searchStr string) ([]*entities.AdsData, error) {
	adsDatas, err := service.repo.AdsDataRepository.SearchFullText(ctx, searchStr)
	if err != nil {
		return nil, err
	}
	for _, ads := range adsDatas {
		tags, _ := service.repo.TagsRepository.FindByAdsID(ctx, ads.ID)
		ads.TagsObj = tags

		images, _ := service.repo.ImageUrlRepository.FindByAdsID(ctx, ads.ID)
		ads.ImageUrlObj = images
	}

	return adsDatas, nil
}
