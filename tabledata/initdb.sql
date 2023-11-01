CREATE DATABASE paybook;

\c paybook;

CREATE TABLE users(
    id          BYTEA PRIMARY KEY,
	userimage   TEXT,
	username    TEXT,
	email       TEXT UNIQUE,
	phone       BIGINT,
	gender      TEXT,
	dateofbirth TEXT,
	password    TEXT,
	roles       TEXT,
	created_at   TIMESTAMP WITH TIME ZONE,
	updated_at   TIMESTAMP WITH TIME ZONE,
	deleted_at   TIMESTAMP WITH TIME ZONE
);

CREATE TABLE addresses(
    id             BYTEA PRIMARY KEY,
	user_id         BYTEA,
	address_type    TEXT,
	recipient_name  TEXT,
	recipient_phone TEXT,
	address_name    TEXT,
	postal_code     TEXT,
	city           TEXT,
	created_at      TIMESTAMP WITH TIME ZONE,
	updated_at      TIMESTAMP WITH TIME ZONE,
	deleted_at      TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);