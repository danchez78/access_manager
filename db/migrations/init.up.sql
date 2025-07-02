CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE SCHEMA access_manager;

CREATE TABLE access_manager.users
(
    id VARCHAR(255) NOT NULL,
    user_agent VARCHAR NOT NULL,
    ip_address VARCHAR NOT NULL,
    is_deauthorised BOOL NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE access_manager.refresh_tokens
(
    user_id VARCHAR(255) NOT NULL,
    issued_time int NOT NULL,
    expiration_time int NOT NULL,
    string text NOT NULL,
    PRIMARY KEY (user_id)
);
