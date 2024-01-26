-- +goose Up
-- +goose StatementBegin
create table IF NOT EXISTS events
(
    id            serial primary key,
    name            text NOT NULL,
    description   text      not null,
    date          timestamp not null
);

CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    email text NOT NULL,
    hash TEXT NOT NULL
);

CREATE TABLE if NOT EXISTS viewer
(
    id SERIAL PRIMARY KEY,
    id_event INTEGER REFERENCES EVENTS(id),
    id_user INTEGER REFERENCES USERS(ID)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table IF EXISTS events;
DROP TABLE if EXISTS USERS;
DROP TABLE IF EXISTS VIEWER;
-- +goose StatementEnd
