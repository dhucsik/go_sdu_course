package database

var migration = `
	create table if not exists categories (
	category_id serial not null primary key,
	category_title varchar(100) not null);
	
	create table if not exists users (
	user_id serial not null primary key,
	username varchar(50) not null unique,
	created_at timestamp,
	updated_at timestamp,
	email varchar(255) not null unique,
	password_hash varchar(255) not null,
	user_role varchar(25) not null);
	
	create table if not exists orders (
	order_id serial not null primary key,
	user_id integer not null references users(user_id),
	total_price numeric(10, 2) not null,
	date timestamp not null);
	
	create table if not exists products (
	product_id serial not null primary key,
	product_title varchar(100) not null,
	seller_id integer not null references users(user_id),
	category_id integer not null references categories(category_id),
	price numeric(10, 2) not null,
	description varchar(300) not null,
	avg_rating numeric(3, 2));
	
	create table if not exists ordered_items (
	ordered_item_id serial not null primary key,
	order_id integer not null references orders(order_id),
	product_id integer not null references products(product_id),
	quantity integer not null,
	price numeric(10, 2) not null);
	
	create table if not exists payments (
	payment_id serial not null primary key,
	order_id integer not null references orders(order_id),
	payment_method varchar(20) not null,
	transaction_id varchar(50) not null,
	status varchar(20) not null,
	created_at timestamp not null,
	updated_at timestamp not null);
	
	create table if not exists reviews (
	review_id serial not null primary key,
	user_id integer not null references users(user_id),
	product_id integer not null references products(product_id),
	rating integer not null,
	comment text,
	constraint unique_user_product_review unique (user_id, product_id)
	);
`
