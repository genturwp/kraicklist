package entities

type AdsData struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	ThumbUrl  string     `json:"thumb_url"`
	Tags      []Tag      `json:"tags"`
	UpdatedAt int64      `json:"updated_at"`
	ImageUrls []ImageUrl `json:"image_urls"`
	RowHash   string     `json:"rowHash"`
}

type ImageUrl struct {
	ID    int64  `json:"id"`
	Image string `json:"image"`
}

type Tag struct {
	ID      int64  `json:"id"`
	TagName string `json:"tag_name"`
}
