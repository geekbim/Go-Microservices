CREATE TABLE IF NOT EXISTS users
(
    id         UUID                     NOT NULL UNIQUE,
    app        VARCHAR(1)               NULL DEFAULT 'A',
    name       VARCHAR                  NULL DEFAULT NULL,
    email      VARCHAR                  NULL DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE NULL DEFAULT NULL,
    PRIMARY KEY (id)
);