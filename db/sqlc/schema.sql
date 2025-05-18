CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    password text NOT NULL
);

CREATE INDEX IF NOT EXISTS users_name_index ON users(name);

CREATE TABLE IF NOT EXISTS exchange_rate(
    id bigserial PRIMARY KEY,
    from_currency VARCHAR(255),
    to_currency VARCHAR(255),
    rate real,
    date timestamp with time zone DEFAULT now()
);
CREATE TABLE IF NOT EXISTS article(
    id bigserial PRIMARY KEY,
    title VARCHAR(255),
    content  text,
    preview  text
);