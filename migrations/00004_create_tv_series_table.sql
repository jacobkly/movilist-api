-- +goose Up
-- +goose StatementBegin
CREATE TABLE tv_series (
    tv_id                 INTEGER PRIMARY KEY, -- TMDB series ID
    adult                   BOOLEAN DEFAULT FALSE,
    backdrop_path           TEXT,
    created_by              JSONB,              -- array of {id, credit_id, name, gender, profile_path}
    episode_run_time         INTEGER[],         -- array of runtimes in minutes
    first_air_date          DATE,
    genres                  JSONB,              -- array of {id, name}
    homepage                TEXT,
    in_production           BOOLEAN DEFAULT TRUE,
    languages               TEXT[],             -- array of language codes
    last_air_date           DATE,
    last_episode_to_air     JSONB,              -- {id, name, overview, vote_average, vote_count, air_date, episode_number, production_code, season_number, runtime, still_path, show_id}
    next_episode_to_air     JSONB,              -- same structure as above, nullable
    networks                JSONB,              -- array of {id, logo_path, name, origin_country}
    number_of_episodes      INTEGER DEFAULT 0,
    number_of_seasons       INTEGER DEFAULT 0,
    origin_country          TEXT[],             -- array of country codes
    original_language       TEXT,
    original_name           TEXT,
    overview                TEXT,
    popularity              DOUBLE PRECISION DEFAULT 0,
    poster_path             TEXT,
    production_companies    JSONB,              -- array of {id, logo_path, name, origin_country}
    production_countries    JSONB,              -- array of {iso_3166_1, name}
    seasons                 JSONB,              -- array of {air_date, episode_count, id, name, overview, poster_path, season_number, vote_average}
    spoken_languages        JSONB,              -- array of {english_name, iso_639_1, name}
    status                  TEXT,
    tagline                 TEXT,
    type                    TEXT,
    vote_average            DOUBLE PRECISION DEFAULT 0,
    vote_count              INTEGER DEFAULT 0,
    created_at              TIMESTAMPTZ DEFAULT NOW(),
    updated_at              TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tv_series;
-- +goose StatementEnd
