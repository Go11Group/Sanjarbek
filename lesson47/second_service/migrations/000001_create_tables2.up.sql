CREATE TABLE IF NOT EXISTS weather (
    name VARCHAR NOT NULL,
    temperature INTEGER,
    humidity REAL,
    wind_speed INTEGER,
    condition VARCHAR,
    date DATE DEFAULT CURRENT_DATE
);

CREATE TABLE IF NOT EXISTS transport (
    number INTEGER NOT NULL,
    stations TEXT[],
    current_station VARCHAR(255),
    is_full BOOLEAN
)