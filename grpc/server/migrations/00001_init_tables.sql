-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE library (
                         id serial primary key,
                         name varchar(255) not null,
                         author varchar(255) not null,
                         year varchar(255) not null
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE library;
