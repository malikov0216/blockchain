-- ************************************** "public".users
CREATE TABLE users (
    id serial NOT NULL unique,
    email varchar(255) NOT NULL unique,
    password varchar(255) NOT NULL,
    first_name varchar(255) NOT NULL,
    second_name varchar(255) NOT NULL,
    created_at timestamp without time zone NULL default now(),
    updated_at timestamp without time zone NULL,
    deleted_at timestamp without time zone NULL
);

CREATE TABLE currencies (
    id serial NOT NULL unique,
    "name" varchar(50) NOT NULL,
    created_at timestamp without time zone NULL default now(),
    updated_at timestamp without time zone NULL,
    deleted_at timestamp without time zone NULL
);

CREATE TABLE wallets (
    id serial NOT NULL unique,
    address varchar(255) NOT NULL,
    currency_id int references currencies(id) on delete cascade NOT NULL,
    balance double precision NOT NULL default 100.0,
    user_id int references users(id) on delete cascade NOT NULL,
    created_at timestamp without time zone NULL default now(),
    updated_at timestamp without time zone NULL,
    deleted_at timestamp without time zone NULL
);

CREATE TABLE transactions (
    id serial NOT NULL unique,
    sum int NOT NULL,
    commission int NOT NULL,
    user_id int references users(id) on delete cascade NOT NULL,
    currency_name VARCHAR(50) NOT NULL,
    address_from VARCHAR(255) NOT NULL,
    address_to VARCHAR(255) NOT NULL,
    created_at timestamp without time zone NULL default now(),
    updated_at timestamp without time zone NULL,
    deleted_at timestamp without time zone NULL
);

INSERT INTO currencies (name)
values ('BTC');

INSERT INTO currencies (name)
values ('ETH');