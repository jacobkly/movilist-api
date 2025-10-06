-- +goose Up
-- +goose StatementBegin
CREATE TABLE tv_collections (
    collection_id    SERIAL PRIMARY KEY, -- custom collection ID
    tv_id          INTEGER NOT NULL,   -- TV series ID
    seasons          JSONB,              -- array of {season_number, poster_path, name, position}
    name             TEXT NOT NULL,      -- collection name
    created_at       TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tv_collections;
-- +goose StatementEnd
