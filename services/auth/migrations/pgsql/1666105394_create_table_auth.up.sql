CREATE TABLE IF NOT EXISTS auth
(
    id         UUID                     NOT NULL UNIQUE,
    app        VARCHAR(1)               NULL DEFAULT 'A',
    user_id    UUID                     NOT NULL UNIQUE,
    email      VARCHAR                  NULL DEFAULT NULL,
    password   VARCHAR                  NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE NULL DEFAULT NULL,
    PRIMARY KEY (id)
);