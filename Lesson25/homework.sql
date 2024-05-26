-- Create the brand table
CREATE TABLE brand (
    id    uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name  varchar          NOT NULL,
    year  int              NOT NULL
);

-- Create the car table
CREATE TABLE car (
    id       uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name     varchar          NOT NULL,
    year     int              NOT NULL,
    colour   varchar          NOT NULL DEFAULT 'black',
    brand_id uuid             NOT NULL REFERENCES brand(id)
);

-- Create the car_brand join table
CREATE TABLE car_brand (
    car_id   uuid NOT NULL REFERENCES car(id),
    brand_id uuid NOT NULL REFERENCES brand(id),
    PRIMARY KEY (car_id, brand_id)
);

