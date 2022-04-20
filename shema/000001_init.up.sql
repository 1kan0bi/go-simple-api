CREATE TABLE products 
(
    id int not null unique,
    name varchar(255) not null,
    price int not null,
    count int,
    description varchar(255)
);


INSERT INTO products VALUES (1, 'IPhone 13 Pro', 130000, 10, 'Smartphone by Apple');
INSERT INTO products VALUES (2, 'IPhone 12 Pro', 120000, 10, 'Smartphone by Apple');
INSERT INTO products VALUES (3, 'IPhone 11 Pro', 110000, 10, 'Smartphone by Apple');
INSERT INTO products VALUES (4, 'IPhone 11', 100000, 10, 'Smartphone by Apple');