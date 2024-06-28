-- +migrate Up
CREATE TABLE Users (
    id int,
    username text
);

-- +migrate Down
DROP TABLE Users;