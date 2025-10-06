-- +goose Up
-- +goose StatementBegin
CREATE TABLE movies (
    movie_id             INTEGER PRIMARY KEY,  -- TMDB movie ID
    adult               BOOLEAN DEFAULT FALSE,
    backdrop_path       TEXT,
    belongs_to_collection JSONB,              -- holds {id, name, poster_path, backdrop_path}
    budget              INTEGER DEFAULT 0,
    genres              JSONB,                -- array of {id, name}
    homepage            TEXT,
    imdb_id             TEXT,
    original_language   TEXT,
    original_title      TEXT,
    overview            TEXT,
    popularity          DOUBLE PRECISION DEFAULT 0,
    poster_path         TEXT,
    production_companies JSONB,               -- array of {id, logo_path, name, origin_country}
    production_countries JSONB,               -- array of {iso_3166_1, name}
    release_date        DATE,
    revenue             BIGINT DEFAULT 0,
    runtime             INTEGER DEFAULT 0,
    spoken_languages    JSONB,                -- array of {english_name, iso_639_1, name}
    status              TEXT,
    tagline             TEXT,
    title               TEXT,
    video               BOOLEAN DEFAULT FALSE,
    vote_average        DOUBLE PRECISION DEFAULT 0,
    vote_count          INTEGER DEFAULT 0,
    created_at          TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at          TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movies;
-- +goose StatementEnd
