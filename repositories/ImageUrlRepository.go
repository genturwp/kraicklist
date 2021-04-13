package repositories

import (
	"context"

	"challenge.haraj.com.sa/kraicklist/entities"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ImageUrlRepository interface {
	Save(ctx context.Context, imageUrl *entities.ImageUrl) (*entities.ImageUrl, error)
	FindByAdsID(ctx context.Context, adsID int64) ([]*entities.ImageUrl, error)
}

type imgUrlRepo struct {
	DB *pgxpool.Pool
}

func NewImageUrlRepository(db *pgxpool.Pool) ImageUrlRepository {
	return &imgUrlRepo{
		DB: db,
	}
}

func (repo *imgUrlRepo) Save(ctx context.Context, imageUrl *entities.ImageUrl) (*entities.ImageUrl, error) {
	query := `
		INSERT INTO image_urls(image, ads_data_id)
		VALUES($1, $2)
		RETURNING id, image, ads_data_id
	`
	var (
		_ID        pgtype.Int8
		_Image     pgtype.Varchar
		_AdsDataId pgtype.Int8
	)

	err := repo.DB.QueryRow(ctx, query, imageUrl.Image, imageUrl.AdsDataID).
		Scan(&_ID, &_Image, &_AdsDataId)
	if err != nil {
		return nil, err
	}

	return &entities.ImageUrl{
		ID:        _ID.Int,
		AdsDataID: _AdsDataId.Int,
		Image:     _Image.String,
	}, nil
}

func (repo *imgUrlRepo) FindByAdsID(ctx context.Context, adsID int64) ([]*entities.ImageUrl, error) {
	query := `
		SELECT id, image, ads_data_id
		FROM image_urls
		WHERE ads_data_id = $1
	`
	var (
		_ID        pgtype.Int8
		_Image     pgtype.Varchar
		_AdsDataId pgtype.Int8
	)
	images, err := repo.DB.Query(ctx, query, adsID)
	if err != nil {
		return nil, err
	}

	result := make([]*entities.ImageUrl, 0)
	for images.Next() {
		err = images.Scan(&_ID, &_Image, &_AdsDataId)
		if err != nil {
			return nil, err
		}
		image := &entities.ImageUrl{
			ID:        _ID.Int,
			AdsDataID: _AdsDataId.Int,
			Image:     _Image.String,
		}
		result = append(result, image)
	}
	return result, nil
}
