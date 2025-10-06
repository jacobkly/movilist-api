-- +goose Up
-- +goose StatementBegin
CREATE TABLE media_index (
    media_id       SERIAL PRIMARY KEY,           -- internal unified ID
    tmdb_id        INTEGER NOT NULL,             -- TMDB movie ID or TV series ID
    season_number  INTEGER,                      -- only for type = 'season'
    media_type     TEXT NOT NULL CHECK (media_type IN ('movie', 'season')), -- seasonal due to season-level nature of MoviList
    created_at     TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(tmdb_id, season_number, media_type)         -- ensures unique mapping for movie or season
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS media_index;
-- +goose StatementEnd
