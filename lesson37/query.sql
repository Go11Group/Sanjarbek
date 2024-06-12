CREATE TABLE users(
    id uuid default gen_random_uuid(),
    first_name varchar,
    last_name varchar,
    age int,
    gender varchar,
    nation text,
    field varchar,
    parent_name varchar,
    city varchar
)