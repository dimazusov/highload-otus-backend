-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users
(
    id INT(11) NOT NULL AUTO_INCREMENT,
    password VARCHAR(32) NOT NULL,
    name VARCHAR(32) NOT NULL,
    surname VARCHAR(32) NOT NULL,
    age INT(11) NOT NULL,
    sex BOOLEAN NOT NULL,
    city VARCHAR(32) NOT NULL,
    intersect VARCHAR(32) NOT NULL,
    CONSTRAINT id_pk PRIMARY KEY (id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table users;