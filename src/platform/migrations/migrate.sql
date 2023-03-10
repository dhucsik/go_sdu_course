create table categories (
id serial not null primary key,
title varchar(100) not null);

/*
CREATE TABLE
e_commerce_db=# \d
                List of relations
 Schema |       Name        |   Type   |  Owner   
--------+-------------------+----------+----------
 public | categories        | table    | postgres
 public | categories_id_seq | sequence | postgres
(2 rows)
*/

create table users (
id serial not null primary key,
username varchar(50) not null unique,
created_at timestamp,
updated_at timestamp,
email varchar(255) not null unique,
password_hash varchar(255) not null,
user_role varchar(25) not null);

/*
CREATE TABLE
e_commerce_db=# \d
                List of relations
 Schema |       Name        |   Type   |  Owner   
--------+-------------------+----------+----------
 public | categories        | table    | postgres
 public | categories_id_seq | sequence | postgres
 public | users             | table    | postgres
 public | users_id_seq      | sequence | postgres
(4 rows)
*/


create table orders (
id serial not null primary key,
user_id integer not null references users(id),
total_price numeric(10, 2) not null,
date timestamp not null);

create table products (
id serial not null primary key,
title varchar(100) not null,
seller_id integer not null references users(id),
category_id integer not null references categories(id),
price numeric(10, 2) not null,
description varchar(300) not null);

create table ordered_items (
id serial not null primary key,
order_id integer not null references orders(id),
product_id integer not null references products(id),
quantity integer not null,
price numeric(10, 2) not null);

create table payments (
id serial not null primary key,
order_id integer not null references orders(id),
payment_method varchar(20) not null,
transaction_id varchar(50) not null,
status varchar(20) not null,
created_at timestamp not null,
updated_at timestamp not null);

