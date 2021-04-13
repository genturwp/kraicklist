package repositories

import (
	"context"

	"challenge.haraj.com.sa/kraicklist/entities"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TagsRepository interface {
	Save(ctx context.Context, tag *entities.Tag) (*entities.Tag, error)
	FindByAdsID(ctx context.Context, adsID int64) ([]*entities.Tag, error)
}

type tagsRepo struct {
	DB *pgxpool.Pool
}

func NewTagsRepository(db *pgxpool.Pool) TagsRepository {
	return &tagsRepo{
		DB: db,
	}
}

func (repo *tagsRepo) Save(ctx context.Context, tag *entities.Tag) (*entities.Tag, error) {
	query := `
		INSERT INTO tags(tag_name, ads_data_id)
		VALUES($1, $2)
		RETURNING id, tag_name, ads_data_id
	`
	var (
		_ID        pgtype.Int8
		_TagName   pgtype.Varchar
		_AdsDataId pgtype.Int8
	)
	err := repo.DB.QueryRow(ctx, query, tag.TagName, tag.AdsDataID).
		Scan(&_ID, &_TagName, &_AdsDataId)
	if err != nil {
		return nil, err
	}

	return &entities.Tag{
		ID:        _ID.Int,
		AdsDataID: _AdsDataId.Int,
		TagName:   _TagName.String,
	}, nil
}

func (repo *tagsRepo) FindByAdsID(ctx context.Context, adsID int64) ([]*entities.Tag, error) {
	query := `
		SELECT id, tag_name, ads_data_id
		FROM tags
		WHERE ads_data_id = $1
	`
	var (
		_ID        pgtype.Int8
		_TagName   pgtype.Varchar
		_AdsDataId pgtype.Int8
	)
	tags, err := repo.DB.Query(ctx, query, adsID)
	if err != nil {
		return nil, err
	}
	result := make([]*entities.Tag, 0)
	for tags.Next() {
		err = tags.Scan(&_ID, &_TagName, &_AdsDataId)
		if err != nil {
			return nil, err
		}
		tag := &entities.Tag{
			ID:        _ID.Int,
			AdsDataID: _AdsDataId.Int,
			TagName:   _TagName.String,
		}
		result = append(result, tag)
	}
	return result, nil
}
