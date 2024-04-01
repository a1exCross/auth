-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD username TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN username;
-- +goose StatementEnd
