-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE library
ADD created_at timestamp not null default now(),
ADD updated_at timestamp not null default now();

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.