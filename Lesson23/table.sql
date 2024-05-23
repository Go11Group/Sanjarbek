CREATE TABLE students (id SERIAL PRIMARY KEY, name VARCHAR(50), stipend FLOAT4, course INT);

INSERT INTO students (name, stipend, course) VALUES
('Sanjarbek', 300000, 4),
('Fahriddin', 500000, 2),
('Begzod', 4000000, 2),
('Diyorbek', 300000, 3),
('Hamidjon', 400000, 3),
('Nurmuhammad', 350000, 1),
('Ali', 200000, 1),
('Vali', 450000, 4),
('Eshmat', 250000, 2),
('Toshmat', 150000, 1);


CREATE TABLE courses (course_id INT PRIMARY KEY, course_name VARCHAR(100));

INSERT INTO courses (course_id, course_name) VALUES
(1, 'Mathematics'),
(2, 'Computer Science'),
(3, 'Physics'),
(4, 'Chemistry');

