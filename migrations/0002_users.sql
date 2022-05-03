-- +goose Up
CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE users.creds (
    "user" TEXT PRIMARY KEY,
    "hash" TEXT
);