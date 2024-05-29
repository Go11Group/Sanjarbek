CREATE TABLE brand (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    name VARCHAR(50),
    year INT
);

CREATE TABLE car (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    brand_id UUID REFERENCES brand(id) NOT NULL,
    name TEXT,
    year INT,
    price NUMERIC
);
