-- +goose Up
CREATE TABLE products
(
    id    serial       not null unique,
    name  varchar(255) not null unique,
    price integer      not null default 0
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE products;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
