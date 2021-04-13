package services

import "challenge.haraj.com.sa/kraicklist/repositories"

type Service struct {
	AdsDataService AdsDataService
}

func NewService(repo *repositories.Repository) *Service {
	adsDataService := NewAdsDataService(repo)
	return &Service{
		AdsDataService: adsDataService,
	}
}
