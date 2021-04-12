CREATE TABLE IF NOT EXISTS ads_datas (
    id BIGSERIAL NOT NULL,
    title VARCHAR(2000) NOT NULL,
    content TEXT NULL,
    thumb_url VARCHAR(2000),
    updated_at bigint,
    CONSTRAINT ads_datas_pk PRIMARY KEY (id)
)