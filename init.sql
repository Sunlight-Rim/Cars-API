-- CREATE SCHEMA

CREATE SCHEMA IF NOT EXISTS api;

-- DROP TABLES

DROP TABLE IF EXISTS api.users;
DROP TABLE IF EXISTS api.cars;

-- CREATE TABLES

CREATE TABLE IF NOT EXISTS api.users(
    id             SERIAL PRIMARY KEY,
    username       VARCHAR(200) NOT NULL,
    email          VARCHAR(200) UNIQUE NOT NULL,
    phone          BIGINT NOT NULL,
    password_hash  BYTEA NOT NULL,
    deleted        BOOLEAN DEFAULT FALSE NOT NULL,
    created_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS api.cars(
    id             SERIAL PRIMARY KEY,
    user_id        INTEGER NOT NULL REFERENCES api.users(id),
    plate          VARCHAR(6) UNIQUE NOT NULL,
    model          VARCHAR(500) NOT NULL,
    color          VARCHAR(200) NOT NULL,
    created_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- FILL TABLES

INSERT INTO api.users (
    username,
    email,
    phone,
    password_hash
) VALUES (
    'johndoe',
    'me@mail.com',
    83641845273,
    E'\\x5E884898DA28047151D0E56F8DC6292773603D0D6AABBDD62A11EF721D1542D8' -- "password"
);

INSERT INTO api.cars (
    user_id, plate,  model,  color
) VALUES
    (1, 'aaa000', 'Tesla X', 'Grey'),
    (1, 'aaa123', 'BWM X3', 'Yellow'),
    (1, 'aaa333', 'McLaren P1', 'Pink');