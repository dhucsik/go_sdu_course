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
email varchar(255) not null unique,
password varchar(50) not null);

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

create table buyers (
user_id integer not null primary key references users(id),
full_name varchar(100) not null,
billing_address varchar(255) not null,
phone_number varchar(20) not null);

/*
CREATE TABLE
e_commerce_db=# \d buyers
                           Table "public.buyers"
     Column      |          Type          | Collation | Nullable | Default 
-----------------+------------------------+-----------+----------+---------
 user_id         | integer                |           | not null | 
 full_name       | character varying(100) |           | not null | 
 billing_address | character varying(255) |           | not null | 
 phone_number    | character varying(20)  |           | not null | 
Indexes:
    "buyers_pkey" PRIMARY KEY, btree (user_id)
Foreign-key constraints:
    "buyers_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id)
*/


create table sellers (
user_id integer not null primary key references users(id),
full_name varchar(100) not null,
business_address varchar(255) not null,
passport_id varchar(50) not null,
phone_number varchar(20) not null);

/*
CREATE TABLE
e_commerce_db=# \d sellers
                           Table "public.sellers"
      Column      |          Type          | Collation | Nullable | Default 
------------------+------------------------+-----------+----------+---------
 user_id          | integer                |           | not null | 
 full_name        | character varying(100) |           | not null | 
 business_address | character varying(255) |           | not null | 
 passport_id      | character varying(50)  |           | not null | 
 phone_number     | character varying(20)  |           | not null | 
Indexes:
    "sellers_pkey" PRIMARY KEY, btree (user_id)
Foreign-key constraints:
    "sellers_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id)
*/

create table orders (
id serial not null primary key,
buyer_id integer not null references buyers(user_id),
seller_id integer not null references sellers(user_id),
total_price numeric(10, 2) not null,
date timestamp not null);

create table products (
id serial not null primary key,
title varchar(100) not null,
seller_id integer not null references sellers(user_id),
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

