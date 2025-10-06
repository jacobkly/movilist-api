-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id                UUID PRIMARY KEY,                              -- matches Supabase auth UID
    username          TEXT UNIQUE NOT NULL,                          -- display name
    bio               TEXT,                                          -- user bio
    avatar_url        TEXT,                                          -- profile picture URL
    banner_url        TEXT,                                          -- banner / cover image
    favorite_media    JSONB DEFAULT '[]'::JSONB,                     -- array of media objects: {
                                                                      --     media_id: INT,
                                                                      --     media_type: TEXT ('movie' or 'season'),
                                                                      --     season_number: INT,        -- optional, only for 'season'
                                                                      --     poster_path: TEXT,
                                                                      --     name: TEXT,
                                                                      --     year: INT
                                                                      --}
    favorite_people   JSONB DEFAULT '[]'::JSONB,                     -- array of people objects {id, name, profile_path}
    is_active         BOOLEAN DEFAULT FALSE,
    created_at        TIMESTAMPTZ DEFAULT NOW(),
    updated_at        TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
