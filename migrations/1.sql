-- +goose Up
CREATE SCHEMA IF NOT EXISTS common;

CREATE TABLE common.projects (
    id   SERIAL,
    name TEXT,
    link CHARACTER VARYING(10)
);

-- +goose Down
DROP TABLE common.projects;
DROP SCHEMA IF EXISTS common;
