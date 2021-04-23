DROP TABLE IF EXISTS products;

CREATE TABLE products (
  id UUID NOT NULL PRIMARY KEY,
  price_in_cents int NOT NULL,
  title varchar(100) NOT NULL,
  description varchar(255) DEFAULT NULL
);

INSERT INTO products (id, price_in_cents, title, description)
values 
	('f0b19f39-b0a8-4b39-a42e-6d83f5cbd4aa', 470000,'Playsation 5','Sony video game'),
	('827bd2cc-9537-4330-ad2b-40885344b71a', 500000,'Xbox One X','Microsoft video game'),
	('8bf94cd1-98db-44dd-b82a-ce7af028b677', 200000,'Switch','Nintendo video game'),
	('3d2d1916-e11c-4b87-a514-97b082632b0b', 500000,'PC','A Personal Computer');


DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id UUID NOT NULL PRIMARY KEY,
  first_name varchar(50) NOT NULL,
  last_name varchar(50) NOT NULL,
  date_of_birth date NOT NULL
);

INSERT INTO users (id, first_name, last_name, date_of_birth)
VALUES
	('9f4abbe9-113c-4099-8bef-2184d030f2c3','Felipe','Cardozo', now()),
	('96079cd0-73a7-480a-a8bb-b6bffa0aedb7','Foo','Bar','2001-11-22'),
	('2b6719be-f4b8-4bc4-8989-13f0c4d0d9b8','Bar','Foo','1993-04-10');