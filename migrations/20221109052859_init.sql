-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE users (
    id          serial primary key,
    name        varchar(128) NOT NULL,
    login       varchar(128) NOT NULL,
    password    varchar(61) NOT NULL,
    created_at  timestamp NOT NULL,
    updated_at  timestamp NOT NULL
);

INSERT INTO users 
        (name, login, password, created_at, updated_at)
    VALUES
        ('admin', 'admin', '$2a$10$GITM6eEDIbH2jn0u6skbae8SqDUI10OFNIJWNBj9LeSRcavbhub/G', current_timestamp, current_timestamp);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE users;

-- +goose StatementEnd
