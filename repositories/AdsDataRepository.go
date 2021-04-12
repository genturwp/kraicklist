package repositories

import (
	"context"

	"challenge.haraj.com.sa/kraicklist/entities"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AdsDataRepository interface {
	Save(ctx context.Context, adsData *entities.AdsData) (*entities.AdsData, error)
	SearchFullText(ctx context.Context, searchStr string) ([]*entities.AdsData, error)
}

type repo struct {
	DB *pgxpool.Pool
}

func NewAdsDataRepository(db *pgxpool.Pool) AdsDataRepository {
	return &repo{
		DB: db,
	}
}

func (repo *repo) Save(ctx context.Context, adsData *entities.AdsData) (*entities.AdsData, error) {
	query := `
		INSERT INTO ads_data(title, content, thumb_url, updated_at, row_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, title, content, thumb_url, updated_at, row_hash
	`
	var (
		_ID        pgtype.Int8
		_Title     pgtype.Varchar
		_Content   pgtype.Varchar
		_ThumbUrl  pgtype.Varchar
		_UpdatedAt pgtype.Int8
		_RowHash   pgtype.Varchar
	)

	err := repo.DB.QueryRow(ctx, query, adsData.Title, adsData.Content, adsData.ThumbUrl, adsData.UpdatedAt, adsData.RowHash).
		Scan(&_ID, &_Title, &_Content, &_ThumbUrl, &_UpdatedAt, &_RowHash)
	if err != nil {
		return nil, err
	}

	return &entities.AdsData{
		ID:        _ID.Int,
		Title:     _Title.String,
		Content:   _Content.String,
		ThumbUrl:  _ThumbUrl.String,
		UpdatedAt: _UpdatedAt.Int,
		RowHash:   _RowHash.String,
	}, nil
}

func (repo *repo) SearchFullText(ctx context.Context, searchStr string) ([]*entities.AdsData, error) {
	return nil, nil
}
