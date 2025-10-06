-- +goose Up
-- +goose StatementBegin
ALTER TABLE media_index
RENAME COLUMN tmdb_id TO id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE media_index
RENAME COLUMN id TO tmdb_id;
-- +goose StatementEnd
