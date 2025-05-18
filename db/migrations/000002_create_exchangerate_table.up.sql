CREATE TABLE IF NOT EXISTS exchange_rate(
    id bigserial PRIMARY KEY,
    from_currency VARCHAR(255),
    to_currency VARCHAR(255),
    rate real,
    date timestamp with time zone DEFAULT now()
);