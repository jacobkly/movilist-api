-- +goose Up
-- +goose StatementBegin
CREATE TABLE media_list (
    id               SERIAL PRIMARY KEY,
    user_id          UUID REFERENCES users(id),
    media_id         INT REFERENCES media_index(media_id),
    media_type             TEXT NOT NULL CHECK (media_type IN ('movie', 'season')),       -- source type for filtering
    status           TEXT CHECK (status IN ('watching', 'planning', 'completed', 'rewatching', 'paused', 'dropped')),
    score            INT,                           -- standardized 0-100%
    episode_progress INT DEFAULT 0,                -- number of episodes watched (for seasons)
    start_date       DATE,
    finish_date      DATE,
    total_rewatches  INT DEFAULT 0,
    notes            TEXT,
    is_deleted       BOOLEAN DEFAULT FALSE,
    created_at       TIMESTAMPTZ DEFAULT NOW(),
    updated_at       TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(user_id, media_id)                     -- ensures no duplicates per user
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS media_list;
-- +goose StatementEnd
