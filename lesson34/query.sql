-- Users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);

INSERT INTO users (username, email, password) VALUES
('johndoe', 'johndoe@example.com', 'johnspassword'),
('janedoe', 'janedoe@example.com', 'janespassword'),
('alice', 'alice@example.com', 'alicepassword'),
('bobsmith', 'bobsmith@example.com', 'bobspassword'),
('charliebrown', 'charliebrown@example.com', 'charliespassword');


-- Products
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT NOT NULL
);

INSERT INTO products (name, description, price, stock_quantity) VALUES
('Laptop', 'A high-performance laptop for everyday use', 999.99, 50),
('Smartphone', 'A latest model smartphone with advanced features', 699.99, 200),
('Headphones', 'Noise-cancelling over-ear headphones', 199.99, 150),
('Smartwatch', 'A smartwatch with fitness tracking capabilities', 249.99, 120),
('Tablet', 'A tablet with a high-resolution display and large storage', 399.99, 80);


-- UserProduct
CREATE TABLE user_products (
    id SERIAL PRIMARY KEY,
    user_id int REFERENCES users(id),
    product_id int REFERENCES products(id)
);

INSERT INTO user_products (user_id, product_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(1, 3),
(2, 5),
(3, 4),
(4, 1),
(5, 2);
