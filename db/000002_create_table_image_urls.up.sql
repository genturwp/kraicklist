CREATE TABLE IF NOT EXISTS image_urls (
    id BIGSERIAL NOT NULL,
    image VARCHAR(2000) NOT NULL,
    CONSTRAINT image_urls_pk PRIMARY KEY(id)
)