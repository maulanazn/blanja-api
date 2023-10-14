CREATE TABLE IF NOT EXISTS paybook(
    id VARCHAR PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id INT NOT NULL UNIQUE,
    title VARCHAR NOT NULL,
    description TEXT NOT NULL,
    writer VARCHAR NOT NULL,
    price INT NOT NULL
);

CREATE INDEX paybok_user_id_key ON paybook(user_id);
ALTER TABLE paybook ADD FOREIGN KEY(user_id) REFERENCES users(id);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL
);

CREATE INDEX users_username_key ON users(username);
ALTER TABLE users ADD UNIQUE(username);

CREATE TABLE IF NOT EXISTS wallet(
    id VARCHAR PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR NOT NULL UNIQUE,
    balance BIGINT NOT NULL,
    card_name VARCHAR NOT NULL
);

CREATE INDEX wallet_username_key ON wallet(username);
ALTER TABLE wallet ADD FOREIGN KEY(username) REFERENCES users(username);