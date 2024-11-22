-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE table if not exists users (
    id serial not null primary key,
    nickname varchar,  -- unique
    email varchar, -- unique
    phone varchar, -- unique
    name varchar,
    surname varchar,
    pwd varchar,
    photo varchar,
    description varchar
);

CREATE table if not exists friends (
    id_user bigint,
    id_friend bigint
);

CREATE table if not exists chat_user (
    id_chat bigint,
    id_user bigint
);

CREATE table if not exists chat (
    id serial not null primary key,
    name varchar,
    description varchar
);

create table if not exists message (
    id serial not null primary key,
    id_chat bigint,
    id_message bigint,
    sent_at timestamp,
    text varchar
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists users, friends, chat, chat_user, message;
-- +goose StatementEnd
