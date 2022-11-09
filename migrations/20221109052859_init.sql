-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE users (
    id          serial primary key,
    name        varchar(128) NOT NULL,
    login       varchar(128) NOT NULL,
    password    varchar(32) NOT NULL,
    created_at  timestamp NOT NULL,
    updated_at  timestamp NOT NULL
);

INSERT INTO users 
        (name, login, password, created_at, updated_at)
    VALUES
        ('admin', 'admin', '123', current_timestamp, current_timestamp);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE users;

-- +goose StatementEnd
