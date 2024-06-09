CREATE table problems(id serial, name varchar, text varchar)

CREATE table users(id serial primary key, first_name varchar, last_name varchar, field varchar, email varchar)

CREATE table solved_problems(id serial primary key, name varchar, degre varchar, user_id int references users(id))