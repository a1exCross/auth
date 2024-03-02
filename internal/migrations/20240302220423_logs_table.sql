-- +goose Up
-- +goose StatementBegin
CREATE TABLE logs (
    id SERIAL PRIMARY KEY,
    action TEXT NOT NULL,
    content TEXT,
    timestamp TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP DATABASE logs;
-- +goose StatementEnd
