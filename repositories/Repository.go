package repositories

import "github.com/jackc/pgx/v4/pgxpool"

type Repository struct {
	AdsDataRepository  AdsDataRepository
	TagsRepository     TagsRepository
	ImageUrlRepository ImageUrlRepository
}

func NewRepository(db *pgxpool.Pool) *Repository {
	adsDataRepository := NewAdsDataRepository(db)
	tagsRepository := NewTagsRepository(db)
	imageUrlRepository := NewImageUrlRepository(db)
	return &Repository{
		AdsDataRepository:  adsDataRepository,
		TagsRepository:     tagsRepository,
		ImageUrlRepository: imageUrlRepository,
	}
}
