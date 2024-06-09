CREATE table problems(id serial, name varchar, difficulty varchar, explanation text)

CREATE table users(id serial primary key, first_name varchar, last_name varchar, field varchar, email varchar)

CREATE table solved_problems(id serial primary key, user_id int references users(id), name varchar, difficulty varchar, explanation text)