package repositories

import (
	"context"
	"strings"

	"challenge.haraj.com.sa/kraicklist/entities"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AdsDataRepository interface {
	Save(ctx context.Context, adsData *entities.AdsData) (*entities.AdsData, error)
	SearchFullText(ctx context.Context, searchStr string) ([]*entities.AdsData, error)
}

type adsDatarepo struct {
	DB *pgxpool.Pool
}

func NewAdsDataRepository(db *pgxpool.Pool) AdsDataRepository {
	return &adsDatarepo{
		DB: db,
	}
}

func (repo *adsDatarepo) Save(ctx context.Context, adsData *entities.AdsData) (*entities.AdsData, error) {
	query := `
		INSERT INTO ads_datas(id, title, content, thumb_url, updated_at, row_hash, doc_vector)
		VALUES ($1, $2, $3, $4, $5, $6, to_tsvector($7))
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
	doc := []string{adsData.Title, adsData.Content}
	docVector := strings.Join(doc, " ")
	err := repo.DB.QueryRow(ctx, query, adsData.ID, adsData.Title, adsData.Content, adsData.ThumbUrl, adsData.UpdatedAt, adsData.RowHash, docVector).
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

func (repo *adsDatarepo) SearchFullText(ctx context.Context, searchStr string) ([]*entities.AdsData, error) {
	query := `SELECT id, title, content, thumb_url, updated_at, row_hash 
	FROM ads_datas
	WHERE doc_vector @@ to_tsquery($1)
	ORDER BY updated_at DESC
	`

	searchWords := strings.ReplaceAll(strings.TrimSpace(searchStr), " ", "|")
	rows, err := repo.DB.Query(ctx, query, searchWords)
	if err != nil {
		return nil, err
	}

	var (
		_ID        pgtype.Int8
		_Title     pgtype.Varchar
		_Content   pgtype.Varchar
		_ThumbUrl  pgtype.Varchar
		_UpdatedAt pgtype.Int8
		_RowHash   pgtype.Varchar
	)

	result := make([]*entities.AdsData, 0)
	for rows.Next() {
		err = rows.Scan(&_ID, &_Title, &_Content, &_ThumbUrl, &_UpdatedAt, &_RowHash)
		if err != nil {
			return nil, err
		}

		adsData := &entities.AdsData{
			ID:        _ID.Int,
			Title:     _Title.String,
			Content:   _Content.String,
			ThumbUrl:  _ThumbUrl.String,
			UpdatedAt: _UpdatedAt.Int,
			RowHash:   _RowHash.String,
		}
		result = append(result, adsData)
	}
	return result, nil
}
