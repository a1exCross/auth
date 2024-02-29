-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id serial primary key,
    name text not null,
    email text,
    role integer,
    password varchar,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP DATABASE users;
-- +goose StatementEnd
