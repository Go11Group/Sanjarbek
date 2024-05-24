CREATE TABLE authors (
    author_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);


CREATE TABLE books (
    book_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    pages INT,
    author_id INT,
    FOREIGN KEY (author_id) REFERENCES authors(author_id) ON DELETE SET NULL
);

INSERT INTO authors (name) VALUES 
('Sanjarbek Jorayev'), 
('Fahriddin Karimov'), 
('Begzod Islomov'), 
('Diyorbek Abdurahmonov'), 
('Hamidjon Usmonov'), 
('Nurmuhammad Ergashev'), 
('Ali Rahmonov'), 
('Vali Toshmatov');

-- Books jadvaliga ma'lumotlar kiritish
INSERT INTO books (name, pages, author_id) VALUES 
('Qahraton', 250, 1), 
('Jazirama', 300, 2), 
('Yolgizlik', 150, 3), 
('Iztirob', 200, 4), 
('Yoshlar Uchun', 220, 5), 
('Daryo Sadosi', 180, 6), 
('Quvonch', 260, 7), 
('Xotiralar', 190, 8), 
('Asrlar Sadosi', 300, 1), 
('Jaholatga Qarshi', 310, 2), 
('Ozodlik', 280, 3), 
('Insoniyat', 270, 4), 
('Kelajak', 240, 5), 
('Tarix', 330, 6), 
('Sarguzasht', 290, 7);
