-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS media
(
    id             UUID PRIMARY KEY,
    title          TEXT      NOT NULL,
    published_date DATE      NOT NULL,
    image_url      TEXT      NULL,
    description    TEXT      NULL,
    created_at     TIMESTAMP NOT NULL,
    updated_at     TIMESTAMP NOT NULL,
    deleted_at     TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS media;
-- +goose StatementEnd
