-- +migrate Up

CREATE TABLE IF NOT EXISTS companies (
    id               VARCHAR(50)     NOT NULL PRIMARY KEY,
    name             VARCHAR(15)     NOT NULL,
    description      TEXT            NOT NULL DEFAULT '',
    employees_amount INTEGER         NOT NULL,
    registered       BOOLEAN         NOT NULL,
    type             VARCHAR(255)    NOT NULL,
    created_at       TIMESTAMP,
    updated_at       TIMESTAMP,

    UNIQUE(name)
);