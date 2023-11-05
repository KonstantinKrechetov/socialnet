CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id            uuid         NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name    VARCHAR(255) NOT NULL,
    second_name   VARCHAR(255) NOT NULL,
    birthdate     timestamp    NOT NULL,
    biography     VARCHAR(511) NOT NULL,
    city          VARCHAR(255) NOT NULL,
    password_hash text         NOT NULL,
    create_time   timestamp    NOT NULL             DEFAULT CURRENT_TIMESTAMP,
    update_time   timestamp    NOT NULL             DEFAULT CURRENT_TIMESTAMP
);
