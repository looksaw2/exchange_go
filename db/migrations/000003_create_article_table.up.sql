CREATE TABLE IF NOT EXISTS article(
    id bigserial PRIMARY KEY,
    title VARCHAR(255),
    content  text,
    preview  text
);