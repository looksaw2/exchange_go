CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    password text NOT NULL
);

CREATE INDEX IF NOT EXISTS users_name_index ON users(name);