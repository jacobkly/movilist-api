-- +goose Up
-- +goose StatementBegin
CREATE TABLE movie_collections (
    collection_id    INTEGER NOT NULL,   -- TMDB collection ID
    movie_id         INTEGER NOT NULL,   -- movie in the collection
    name             TEXT NOT NULL,      -- movie title
    poster_path      TEXT,
    vote_average     DOUBLE PRECISION,
    position         INTEGER,            -- explicit order in the collection
    created_at       TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (collection_id, movie_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movie_collections;
-- +goose StatementEnd
