package entities

type AdsData struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Content     string      `json:"content"`
	ThumbUrl    string      `json:"thumb_url"`
	Tags        []string    `json:"tags"`
	TagsObj     []*Tag      `json:"tags_obj"`
	UpdatedAt   int64       `json:"updated_at"`
	ImageUrlObj []*ImageUrl `json:"image_urls_obj"`
	ImageUrls   []string    `json:"image_urls"`
	RowHash     string      `json:"rowHash"`
}

type ImageUrl struct {
	ID        int64  `json:"id"`
	AdsDataID int64  `json:"ads_data_id"`
	Image     string `json:"image"`
}

type Tag struct {
	ID        int64  `json:"id"`
	AdsDataID int64  `json:"ads_data_id"`
	TagName   string `json:"tag_name"`
}
