CREATE TABLE cars (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    make VARCHAR(255),
    model VARCHAR(255),
    year INT,
    price DECIMAL(10, 2)
);