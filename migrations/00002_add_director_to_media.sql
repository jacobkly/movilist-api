-- +goose Up
-- +goose StatementBegin
ALTER TABLE media ADD COLUMN director TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE media DROP COLUMN director;
-- +goose StatementEnd
