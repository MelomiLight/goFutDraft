CREATE TABLE IF NOT EXISTS players (
    id bigserial PRIMARY KEY,
    common_name text NOT NULL,
    position text NOT NULL,
    league integer NOT NULL,
    nation integer NOT NULL,
    club integer NOT NULL,
    rating integer NOT NULL,
    pace integer NOT NULL,
    shooting integer NOT NULL,
    passing integer NOT NULL,
    dribbling integer NOT NULL,
    defending integer NOT NULL,
    physicality integer NOT NULL
);