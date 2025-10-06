-- +goose Up
-- +goose StatementBegin
CREATE TABLE people (
    person_id           INTEGER PRIMARY KEY,         -- TMDB person ID
    adult               BOOLEAN DEFAULT FALSE,       -- adult content flag
    also_known_as       JSONB DEFAULT '[]'::JSONB,   -- array of alternate names
    biography           TEXT,                        
    birthday            DATE,                       
    deathday            DATE,                    
    gender              INTEGER DEFAULT 0,           -- gender (TMDB coding)
    homepage            TEXT,                      
    imdb_id             TEXT,                     
    known_for_department TEXT,                      
    name                TEXT NOT NULL,            
    place_of_birth      TEXT,                    
    popularity          DOUBLE PRECISION DEFAULT 0,  -- TMDB popularity score
    profile_path        TEXT,                        -- profile image path
    created_at          TIMESTAMPTZ DEFAULT NOW(),
    updated_at          TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS people;
-- +goose StatementEnd
