CREATE TABLE products
(
    id    serial       not null unique,
    name  varchar(255) not null unique,
    price integer      not null default 0
);