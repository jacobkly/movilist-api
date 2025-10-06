-- +goose Up
-- +goose StatementBegin
CREATE TABLE recommendations (
    id              INTEGER NOT NULL,   -- source movie/TV ID
    media_type      TEXT NOT NULL CHECK (media_type IN ('movie', 'season')),      -- source type
    recommended_id  INTEGER NOT NULL,   -- recommended media ID (same type as media_type)
    name            TEXT NOT NULL,      -- title or series name
    poster_path     TEXT,               -- for UI display
    vote_average    DOUBLE PRECISION,   -- optional rating info
    created_at      TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (id, media_type, recommended_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recommendations;
-- +goose StatementEnd
