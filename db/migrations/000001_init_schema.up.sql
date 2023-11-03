CREATE TABLE IF NOT EXISTS users
(
    id            bigserial,
    username      VARCHAR(255) NOT NULL UNIQUE,
    first_name    VARCHAR(255) NOT NULL,
    second_name   VARCHAR(255) NOT NULL,
    birthdate     timestamp    NOT NULL,
    biography     VARCHAR(511) NOT NULL,
    city          VARCHAR(255) NOT NULL,
    password      text         NOT NULL,
    password_salt VARCHAR(255) NOT NULL,
    create_time   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP
);