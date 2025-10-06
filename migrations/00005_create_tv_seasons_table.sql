-- +goose Up
-- +goose StatementBegin
CREATE TABLE tv_seasons (
    tv_id             INTEGER NOT NULL,    -- TMDB series ID
    season_number       INTEGER NOT NULL,    -- Season number within the series
    air_date            DATE,
    name                TEXT,
    overview            TEXT,
    poster_path         TEXT,
    vote_average        DOUBLE PRECISION DEFAULT 0,
    vote_count          INTEGER DEFAULT 0,
    episodes            JSONB,               -- array of episodes, each object: {
                                             --   air_date: string,
                                             --   episode_number: integer,
                                             --   episode_type: string,
                                             --   id: integer,
                                             --   name: string,
                                             --   overview: string,
                                             --   production_code: string,
                                             --   runtime: integer,
                                             --   season_number: integer,
                                             --   show_id: integer,
                                             --   still_path: string,
                                             --   vote_average: number,
                                             --   vote_count: integer,
                                             --   array of crews "crew": {
                                             --       department: string,
                                             --       job: string,
                                             --       credit_id: string,
                                             --       adult: boolean,
                                             --       gender: integer,
                                             --       id: integer,
                                             --       known_for_department: string,
                                             --       name: string,
                                             --       original_name: string,
                                             --       popularity: number,
                                             --       profile_path: string
                                             --   },
                                             --   array of guest stars "guest_stars": {
                                             --       character: string,
                                             --       credit_id: string,
                                             --       order: integer,
                                             --       adult: boolean,
                                             --       gender: integer,
                                             --       id: integer,
                                             --       known_for_department: string,
                                             --       name: string,
                                             --       original_name: string,
                                             --       popularity: number,
                                             --       profile_path: string
                                             --   },
                                             -- }
    networks            JSONB,               -- array of networks: { id: integer, logo_path: string, name: string, origin_country: string }
    created_at          TIMESTAMPTZ DEFAULT NOW(),
    updated_at          TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (tv_id, season_number)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tv_seasons;
-- +goose StatementEnd
