-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS users
(
    id              SERIAL NOT NULL PRIMARY KEY,
    email           varchar UNIQUE,
    name            varchar,
    sur_name        varchar,
    hashed_password VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
