CREATE TABLE IF NOT EXISTS clubs (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    league integer NOT NULL
);