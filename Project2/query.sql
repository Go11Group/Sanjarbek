CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    birthday TIMESTAMP NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);


INSERT INTO users (name, email, birthday, password) VALUES 
('Alice Johnson', 'alice.johnson@example.com', '1985-02-15', 'password1'),
('Bob Smith', 'bob.smith@example.com', '1990-06-25', 'password2'),
('Charlie Brown', 'charlie.brown@example.com', '1982-11-30', 'password3'),
('Diana Prince', 'diana.prince@example.com', '1993-04-18', 'password4'),
('Ethan Hunt', 'ethan.hunt@example.com', '1987-12-12', 'password5'),
('Fiona Apple', 'fiona.apple@example.com', '1995-08-22', 'password6'),
('George Martin', 'george.martin@example.com', '1991-09-09', 'password7'),
('Hannah Montana', 'hannah.montana@example.com', '1996-05-14', 'password8'),
('Ian Malcolm', 'ian.malcolm@example.com', '1988-03-05', 'password9'),
('Jack Sparrow', 'jack.sparrow@example.com', '1981-07-19', 'password10');


CREATE TABLE courses (
    course_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);


INSERT INTO courses (title, description) VALUES 
('Introduction to Programming', 'Learn the basics of programming using Python.'),
('Advanced Data Structures', 'Explore complex data structures and their applications.'),
('Database Systems', 'Understand the principles of database management systems.'),
('Web Development', 'Build dynamic websites using HTML, CSS, and JavaScript.'),
('Machine Learning', 'An introduction to machine learning algorithms and techniques.'),
('Cloud Computing', 'Learn about cloud services and how to deploy applications.'),
('Cybersecurity', 'Understand the fundamentals of cybersecurity and how to protect systems.'),
('Artificial Intelligence', 'Explore the concepts and applications of AI.'),
('Software Engineering', 'Learn about software development methodologies and best practices.'),
('Data Science', 'An overview of data science principles and techniques.');



CREATE TABLE lessons (
    lesson_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0,
    CONSTRAINT fk_course
        FOREIGN KEY(course_id) 
        REFERENCES Courses(course_id)
);


INSERT INTO lessons (lesson_id, course_id, title, content) VALUES 
(gen_random_uuid(), (SELECT course_id FROM courses WHERE title = 'Introduction to Programming'), 'Lesson 1: Introduction to Python', 'Content for Lesson 1'),
(gen_random_uuid(), (SELECT course_id FROM courses WHERE title = 'Introduction to Programming'), 'Lesson 2: Data Types', 'Content for Lesson 2'),
(gen_random_uuid(), (SELECT course_id FROM courses WHERE title = 'Introduction to Programming'), 'Lesson 3: Control Structures', 'Content for Lesson 3'),
(gen_random_uuid(), (SELECT course_id FROM courses WHERE title = 'Advanced Data Structures'), 'Lesson 1: Arrays', 'Content for Lesson 1'),
(gen_random_uuid(), (SELECT course_id FROM courses WHERE title = 'Advanced Data Structures'), 'Lesson 2: Linked Lists', 'Content for Lesson 2'),
(gen_random_uuid(), (SELECT course_id FROM courses WHERE title = 'Database Systems'), 'Lesson 1: Introduction to Databases', 'Content for Lesson 1'),
(gen_random_uuid(), (SELECT course_id FROM courses WHERE title = 'Database Systems'), 'Lesson 2: SQL Basics', 'Content for Lesson 2'),
(gen_random_uuid(), (SELECT course_id FROM courses WHERE title = 'Web Development'), 'Lesson 1: HTML Basics', 'Content for Lesson 1'),
(gen_random_uuid(), (SELECT course_id FROM courses WHERE title = 'Web Development'), 'Lesson 2: CSS Fundamentals', 'Content for Lesson 2'),
(gen_random_uuid(), (SELECT course_id FROM courses WHERE title = 'Machine Learning'), 'Lesson 1: Introduction to ML', 'Content for Lesson 1');


CREATE TABLE enrollments (
    enrollment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    course_id UUID NOT NULL,
    enrollment_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id) 
        REFERENCES Users(user_id),
    CONSTRAINT fk_course
        FOREIGN KEY(course_id) 
        REFERENCES Courses(course_id)
);


INSERT INTO enrollments (enrollment_id, user_id, course_id) VALUES 
(gen_random_uuid(), (SELECT user_id FROM users WHERE email = 'alice.johnson@example.com'), (SELECT course_id FROM courses WHERE title = 'Introduction to Programming')),
(gen_random_uuid(), (SELECT user_id FROM users WHERE email = 'bob.smith@example.com'), (SELECT course_id FROM courses WHERE title = 'Advanced Data Structures')),
(gen_random_uuid(), (SELECT user_id FROM users WHERE email = 'charlie.brown@example.com'), (SELECT course_id FROM courses WHERE title = 'Database Systems')),
(gen_random_uuid(), (SELECT user_id FROM users WHERE email = 'diana.prince@example.com'), (SELECT course_id FROM courses WHERE title = 'Web Development')),
(gen_random_uuid(), (SELECT user_id FROM users WHERE email = 'ethan.hunt@example.com'), (SELECT course_id FROM courses WHERE title = 'Machine Learning')),
(gen_random_uuid(), (SELECT user_id FROM users WHERE email = 'fiona.apple@example.com'), (SELECT course_id FROM courses WHERE title = 'Cloud Computing')),
(gen_random_uuid(), (SELECT user_id FROM users WHERE email = 'george.martin@example.com'), (SELECT course_id FROM courses WHERE title = 'Cybersecurity')),
(gen_random_uuid(), (SELECT user_id FROM users WHERE email = 'hannah.montana@example.com'), (SELECT course_id FROM courses WHERE title = 'Artificial Intelligence')),
(gen_random_uuid(), (SELECT user_id FROM users WHERE email = 'ian.malcolm@example.com'), (SELECT course_id FROM courses WHERE title = 'Software Engineering')),
(gen_random_uuid(), (SELECT user_id FROM users WHERE email = 'jack.sparrow@example.com'), (SELECT course_id FROM courses WHERE title = 'Data Science'));
