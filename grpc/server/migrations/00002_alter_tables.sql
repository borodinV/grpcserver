-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE library
ADD created_at date,
ADD updated_at date;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.